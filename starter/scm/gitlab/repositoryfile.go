package gitlab

import (
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/hidevopsio/hioak/starter/scm"
	"github.com/xanzy/go-gitlab"
)

type RepositoryFile struct {
	scm.TreeNode
	client NewClient
}

func NewRepositoryFile(c NewClient) *RepositoryFile {
	return &RepositoryFile{
		client: c,
	}
}

func (r *RepositoryFile) GetRepository(baseUrl, token, filePath, ref string, pid int) (string, error) {
	log.Debug("Repository.Repository()")
	log.Debugf("url: %v", baseUrl)
	opt := &gitlab.GetFileOptions{
		Ref:      &ref,
		FilePath: &filePath,
	}
	file, _, err := r.client(baseUrl, token).RepositoryFile().GetFile(pid, opt)
	if err != nil {
		return "", err
	}
	return file.Content, nil
}
