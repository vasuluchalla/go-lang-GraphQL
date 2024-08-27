package database

import (
	"context"
	"fmt"
	"go-lang-GraphQL/graph/model"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
)

var connectionString string = "postgres://postgres:password@localhost:5432/go-movies"

type DB struct {
	conn *pgx.Conn
}

// Connect to PostgreSQL database
func Connect() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}

	return &DB{
		conn: conn,
	}
}

// GetJob retrieves a job listing by ID
func (db *DB) GetJob(id string) *model.JobListing {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := `SELECT id, title, description, url, company FROM jobs WHERE id=$1`
	var jobListing model.JobListing

	err := db.conn.QueryRow(ctx, query, id).Scan(&jobListing.ID, &jobListing.Title, &jobListing.Description, &jobListing.URL, &jobListing.Company)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to retrieve job: %v\n", err))
	}
	return &jobListing
}

// GetJobs retrieves all job listings
func (db *DB) GetJobs() []*model.JobListing {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := `SELECT id, title, description, url, company FROM jobs`
	rows, err := db.conn.Query(ctx, query)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to retrieve jobs: %v\n", err))
	}
	defer rows.Close()

	var jobListings []*model.JobListing
	for rows.Next() {
		var jobListing model.JobListing
		err := rows.Scan(&jobListing.ID, &jobListing.Title, &jobListing.Description, &jobListing.URL, &jobListing.Company)
		if err != nil {
			log.Fatal(fmt.Sprintf("Unable to scan job: %v\n", err))
		}
		jobListings = append(jobListings, &jobListing)
	}
	return jobListings
}

// CreateJobListing creates a new job listing
func (db *DB) CreateJobListing(jobInfo model.CreateJobListingInput) *model.JobListing {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := `INSERT INTO jobs (title, description, url, company) VALUES ($1, $2, $3, $4) RETURNING id`
	var insertedID string

	err := db.conn.QueryRow(ctx, query, jobInfo.Title, jobInfo.Description, jobInfo.URL, jobInfo.Company).Scan(&insertedID)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to create job listing: %v\n", err))
	}

	returnJobListing := model.JobListing{
		ID:          insertedID,
		Title:       jobInfo.Title,
		Description: jobInfo.Description,
		URL:         jobInfo.URL,
		Company:     jobInfo.Company,
	}
	return &returnJobListing
}

// UpdateJobListing updates an existing job listing
func (db *DB) UpdateJobListing(jobId string, jobInfo model.UpdateJobListingInput) *model.JobListing {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := `UPDATE jobs SET title=$1, description=$2, url=$3, company=$4 WHERE id=$5 RETURNING id, title, description, url, company`

	// Set up values for update
	title := jobInfo.Title
	description := jobInfo.Description
	url := jobInfo.URL
	company := jobInfo.Company

	// Execute the query
	var jobListing model.JobListing
	err := db.conn.QueryRow(ctx, query, title, description, url, company, jobId).Scan(&jobListing.ID, &jobListing.Title, &jobListing.Description, &jobListing.URL, &jobListing.Company)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to update job listing: %v\n", err))
	}

	return &jobListing
}

// DeleteJobListing deletes a job listing by ID
func (db *DB) DeleteJobListing(jobId string) *model.DeleteJobResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := `DELETE FROM jobs WHERE id=$1`

	_, err := db.conn.Exec(ctx, query, jobId)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to delete job listing: %v\n", err))
	}

	return &model.DeleteJobResponse{DeletedJobID: jobId}
}
