package release

const (
	KeyDataID          = "id"          //yes 	The project ID
	KeyDataTagName     = "tag_name"    //yes	The Git tag the release is associated with.
	KeyDataName        = "name"        //no	The release name.
	KeyDataDescription = "description" //no	The description of the release. You can use Markdown.
	KeyDataMilestones  = "milestones"  //no	The title of each milestone to associate with the release. GitLab Premium customers can specify group milestones. To remove all milestones from the release, specify [].
	KeyDataReleasedAt  = "released_at" //no	The date when the release is/was ready. Expected in ISO 8601 format (2019-03-15T08:00:00Z).
	KeyDataRef         = "ref"         //yes, if tag_name doesn’t exist	If a tag specified in tag_name doesn’t exist, the release is created from ref and tagged with tag_name. It can be a commit SHA, another tag name, or a branch name.
)
