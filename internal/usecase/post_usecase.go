package usecase

import (
	"context"

	"github.com/nujabuz/gososial/internal/entity"
)

type postUseCase struct {
	postRepository PostRepository
}

func NewPostUseCase(postRepository PostRepository) PostUseCase {
	return &postUseCase{
		postRepository: postRepository,
	}
}

func (uc *postUseCase) Delete(c context.Context, in *InputDeletePost) error {
	// check post
	post, err := uc.postRepository.FindByID(c, in.ID)
	if err != nil {
		return err
	}

	if err := post.CheckUserID(in.UserID); err != nil {
		return err
	}

	// delete post
	if err := uc.postRepository.Delete(c, post.ID, post.UserID); err != nil {
		return err
	}

	return nil
}

func (uc *postUseCase) Create(c context.Context, in *InputCreatePost) (*OutputPost, error) {
	// make new post
	post, err := entity.NewPost(in.UserID, in.Body)
	if err != nil {
		return nil, err
	}

	// create post
	if err := uc.postRepository.Create(c, post); err != nil {
		return nil, err
	}

	out := uc.postToOutputPost(post)

	return out, nil
}

func (uc *postUseCase) Update(c context.Context, in *InputUpdatePost) (*OutputPost, error) {
	// check post
	post, err := uc.postRepository.FindByUserID(c, in.UserID)
	if err != nil {
		return nil, err
	}

	// check value for update value objects
	if in.Body != nil {
		if err := post.UpdateBody(*in.Body); err != nil {
			return nil, err
		}
	}

	// update post
	if err := uc.postRepository.Update(c, post); err != nil {
		return nil, err
	}

	// change Post to OutputPost
	out := uc.postToOutputPost(post)

	return out, nil
}

func (uc *postUseCase) FindByID(c context.Context, in *InputFindByIDPost) (*OutputPost, error) {
	// find post by id
	post, err := uc.postRepository.FindByID(c, in.ID)
	if err != nil {
		return nil, err
	}

	// change Post to OutputPost
	out := uc.postToOutputPost(post)

	return out, nil
}

func (uc *postUseCase) FindByUserID(c context.Context, in *InputFindByUserIDPost) (*OutputPost, error) {
	// find post by user id
	post, err := uc.postRepository.FindByUserID(c, in.UserID)
	if err != nil {
		return nil, err
	}

	// change Post to OutputPost
	out := uc.postToOutputPost(post)

	return out, nil
}

func (uc *postUseCase) FindManyByUserID(c context.Context, in *InputFindByUserIDPost) ([]*OutputPost, error) {
	// find post by user id
	posts, err := uc.postRepository.FindManyByUserID(c, in.UserID)
	if err != nil {
		return nil, err
	}

	// change Post to OutputPost
	var out []*OutputPost
	for _, post := range posts {
		out = append(out, uc.postToOutputPost(post))
	}

	return out, nil
}

// helper.
func (uc *postUseCase) postToOutputPost(post *entity.Post) *OutputPost {
	return &OutputPost{
		ID:        post.ID,
		UserID:    post.UserID,
		Body:      post.Body,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}
