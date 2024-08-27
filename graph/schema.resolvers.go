package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"go-lang-GraphQL/database"
	"go-lang-GraphQL/graph/model"
	"log"
)

// CreateJobListing is the resolver for the createJobListing field.
func (r *mutationResolver) CreateJobListing(ctx context.Context, input model.CreateJobListingInput) (*model.JobListing, error) {
	job, err := db.CreateJobListing(input)
	if err != nil {
		log.Printf("Error creating job listing: %v", err)
		return nil, err
	}
	return job, nil
}

// UpdateJobListing is the resolver for the updateJobListing field.
func (r *mutationResolver) UpdateJobListing(ctx context.Context, id string, input model.UpdateJobListingInput) (*model.JobListing, error) {
	job, err := db.UpdateJobListing(id, input)
	if err != nil {
		log.Printf("Error updating job listing with ID %s: %v", id, err)
		return nil, err
	}
	return job, nil
}

// DeleteJobListing is the resolver for the deleteJobListing field.
func (r *mutationResolver) DeleteJobListing(ctx context.Context, id string) (*model.DeleteJobResponse, error) {
	response, err := db.DeleteJobListing(id)
	if err != nil {
		log.Printf("Error deleting job listing with ID %s: %v", id, err)
		return nil, err
	}
	return response, nil
}

// Jobs is the resolver for the jobs field.
func (r *queryResolver) Jobs(ctx context.Context) ([]*model.JobListing, error) {
	jobs, err := db.GetJobs()
	if err != nil {
		log.Printf("Error retrieving job listings: %v", err)
		return nil, err
	}
	return jobs, nil
}

// Job is the resolver for the job field.
func (r *queryResolver) Job(ctx context.Context, id string) (*model.JobListing, error) {
	job, err := db.GetJob(id)
	if err != nil {
		log.Printf("Error retrieving job listing with ID %s: %v", id, err)
		return nil, err
	}
	return job, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
var db = database.Connect()
