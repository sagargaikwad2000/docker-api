package models

type Image struct {
	Containers  int        `json:"containers"`
	Created     int        `json:"created"`
	Id          string     `json:"id"`
	Labels      ImageLabel `json:"labels"`
	ParentId    string     `json:"parentId"`
	RepoDigests []string   `json:"repoDigests"`
	RepoTags    []string   `json:"repoTags"`
	SharedSize  int        `json:"sharedSize"`
	Size        int        `json:"size"`
}

type ImageLabel struct {
	OrgOpencontainersImageCreated     string `json:"org.opencontainers.image.created,omitempty"`
	OrgOpencontainersImageDescription string `json:"org.opencontainers.image.description,omitempty"`
	OrgOpencontainersImageLicenses    string `json:"org.opencontainers.image.licenses,omitempty"`
	OrgOpencontainersImageRefName     string `json:"org.opencontainers.image.ref.name,omitempty"`
	OrgOpencontainersImageRevision    string `json:"org.opencontainers.image.revision,omitempty"`
	OrgOpencontainersImageSource      string `json:"org.opencontainers.image.source,omitempty"`
	OrgOpencontainersImageTitle       string `json:"org.opencontainers.image.title,omitempty"`
	OrgOpencontainersImageUrl         string `json:"org.opencontainers.image.url,omitempty"`
	OrgOpencontainersImageVersion     string `json:"org.opencontainers.image.version,omitempty"`
}
