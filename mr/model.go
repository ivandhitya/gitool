package mr

// RespMR is a type for merge request response
type RespMR struct {
	ID          int      `json:"id"`
	IID         int      `json:"iid"`
	ProjectID   int      `json:"project_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Message     []string `json:"message"`
	User        struct {
		CanMerge bool `json:"can_merge"`
	} `json:"user"`
}

// Merge Request Response sample
// {
//     "id": 17187,
//     "iid": 22,
//     "project_id": 683,
//     "title": "testing merge",
//     "description": null,
//     "state": "opened",
//     "created_at": "2021-09-04T15:51:07.071Z",
//     "updated_at": "2021-09-04T15:51:07.071Z",
//     "merged_by": null,
//     "merged_at": null,
//     "closed_by": null,
//     "closed_at": null,
//     "target_branch": "test-api",
//     "source_branch": "test-api-to-merge",
//     "user_notes_count": 0,
//     "upvotes": 0,
//     "downvotes": 0,
//     "author": {
//         "id": 189,
//         "name": "Ivan Adhi Prasetya",
//         "username": "ivan_prasetya",
//         "state": "active",
//         "avatar_url": "http://gitlab.linkaja.com/uploads/-/system/user/avatar/189/avatar.png",
//         "web_url": "http://gitlab.linkaja.com/ivan_prasetya"
//     },
//     "assignees": [],
//     "assignee": null,
//     "reviewers": [],
//     "source_project_id": 683,
//     "target_project_id": 683,
//     "labels": [],
//     "work_in_progress": false,
//     "milestone": null,
//     "merge_when_pipeline_succeeds": false,
//     "merge_status": "checking",
//     "sha": "ef8e5f0fab4e46af6cec5644f4d95c3d85a57b93",
//     "merge_commit_sha": null,
//     "squash_commit_sha": null,
//     "discussion_locked": null,
//     "should_remove_source_branch": null,
//     "force_remove_source_branch": null,
//     "reference": "!22",
//     "references": {
//         "short": "!22",
//         "relative": "!22",
//         "full": "be/karina!22"
//     },
//     "web_url": "http://gitlab.linkaja.com/be/karina/-/merge_requests/22",
//     "time_stats": {
//         "time_estimate": 0,
//         "total_time_spent": 0,
//         "human_time_estimate": null,
//         "human_total_time_spent": null
//     },
//     "squash": false,
//     "task_completion_status": {
//         "count": 0,
//         "completed_count": 0
//     },
//     "has_conflicts": false,
//     "blocking_discussions_resolved": true,
//     "approvals_before_merge": null,
//     "subscribed": true,
//     "changes_count": "1",
//     "latest_build_started_at": null,
//     "latest_build_finished_at": null,
//     "first_deployed_to_production_at": null,
//     "pipeline": null,
//     "head_pipeline": null,
//     "diff_refs": {
//         "base_sha": "9f3369d1d1e792bfdfbe515851bcf0bcb07c73e1",
//         "head_sha": "ef8e5f0fab4e46af6cec5644f4d95c3d85a57b93",
//         "start_sha": "9f3369d1d1e792bfdfbe515851bcf0bcb07c73e1"
//     },
//     "merge_error": null,
//     "user": {
//         "can_merge": true
//     }
// }
