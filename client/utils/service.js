import m from 'mithril'
import Auth from './auth'

const Service = {
    //>>>>>>>>>>>>> roles
    //getRoles: () => get("/api/roles"),
    //createRole: (name) => post("/api/roles", { name: name }),
    //updateRole: (id, name) => put("/api/roles/" + id, { id: id, name: name }),
    //deleteRole: (id) => del("/api/roles/" + id),

    //>>>>>>>>>>>>> pages
    getPage: (id) => get("/api/pages/" + id),
    getPages: () => get("/api/pages"),
    createPage: (page) => post("/api/pages", page),
    updatePage: (id, page) => put("/api/pages/" + id, page),
    deletePage: (id) => del("/api/pages/" + id),

    //>>>>>>>>>>>>> reports
    getSpent: () => get("/api/reports/spent/"),

    //>>>>>>>>>>>>> categories
    getCategory: (id) => get("/api/categories/" + id),
    getCategories: () => get("/api/categories"),
    createCategory: (category) => post("/api/categories", category),
    updateCategory: (id, category) => put("/api/categories/" + id, category),
    deleteCategory: (id) => del("/api/categories/" + id),
    getCategoriesSummary: () => get("/api/categories_summary"),

    //>>>>>>>>>>>>> projects
    getProject: (id) => get("/api/projects/" + id),
    getProjects: () => get("/api/projects"),
    newProject: () => get("/api/new_project/"),
    createProject: (project) => post("/api/projects", project),
    editProject: (id) => get("/api/edit_project/" + id),
    updateProject: (id, project) => put("/api/projects/" + id, project),
    deleteProject: (id) => del("/api/projects/" + id),
    archiveProject: (id, project) => put("/api/archive_project/" + id, project),
    favorProject: (id, project) => put("/api/favor_project/" + id, project),
    getFavoriteProjects: () => get("/api/favorite_projects"),
    getProjectsSummary: () => get("/api/projects_summary"),

    //>>>>>>>>>>>>>>>>>> tasks
    getTask: (id) => get("/api/tasks/" + id),
    getTasks: () => get("/api/tasks"),
    newTask: (project_id) => get("/api/new_task" + ((project_id) ? "?project_id=" + project_id : "")),
    createTask: (task) => post("/api/tasks", task),
    editTask: (id) => get("/api/edit_task/" + id),
    updateTask: (id, task) => put("/api/tasks/" + id, task),
    deleteTask: (id) => del("/api/tasks/" + id),
    getTasksSummary: () => get("/api/tasks_summary"),

    //>>>>>>>>>>>>>>>>>> comments
    getComment: (id) => get("/api/comments/" + id),
    getComments: (task_id) => get("/api/comments?task_id=" + task_id),
    createComment: (comment) => post("/api/comments", comment),
    updateComment: (id, comment) => put("/api/comments/" + id, comment),
    deleteComment: (id) => del("/api/comments/" + id),

    //>>>>>>>>>>>>>>>>>> sessions
    getSession: (id) => get("/api/sessions/" + id),
    getSessions: () => get("/api/sessions"),
    newSession: () => get("/api/new_session/"),
    createSession: (session) => post("/api/sessions", session),
    deleteSession: (id) => del("/api/sessions/" + id),
    getSessionsSummary: () => get("/api/sessions_summary"),

    //>>>>>>>>>>>>>>>>>> task logs
    getTaskLog: (id) => get("/api/task_logs/" + id),
    getTaskLogs: () => get("/api/task_logs"),
    createTaskLog: (taskLog) => post("/api/task_logs", taskLog),
    updateTaskLog: (id, taskLog) => put("/api/task_logs/" + id, taskLog),

    //>>>>>>>>>>>> notifications
    getNotifications: () => get("/api/notifications"),
    deleteNotification: (id) => del("/api/notifications/" + id),

    //>>>>>>>>>>>>> users
    getUsers: () => get("/api/users"),
    //getProjectUsers: (project_id) => get("/api/project_users/" + project_id),
    getRoles: () => get("/api/roles"),
    getUsersSummary: () => get("/api/users_summary"),

    //>>>>>>>>>>>>>> settings
    getSetting: (id) => get("/api/settings/" + id),
    getSettings: () => get("/api/settings"),
    createSetting: (setting) => post("/api/settings", setting),
    updateSetting: (id, setting) => put("/api/settings/" + id, setting),
    deleteSetting: (id) => del("/api/settings/" + id),

    //>>>>>>>>>>>> upload file
    uploadFile: (file) => {
        let data = new FormData()
        data.append("upload", file)
        return post("/api/upload/form", data);
    },

    //>>>>>>>>>>>>> search
    getSearch: (query) => get("/api/search?query=" + query),

}

const get = (url) => m.request({
    method: "GET",
    url: url,
    headers: {
        Authorization: Auth.authHeader()
    }
})

const post = (url, body) => m.request({
    method: "POST",
    url: url,
    body: body,
    headers: {
        Authorization: Auth.authHeader()
    }
})

const put = (url, body) => m.request({
    method: "PUT",
    url: url,
    body: body,
    headers: {
        Authorization: Auth.authHeader()
    }
})

const del = (url) => m.request({
    method: "DELETE",
    url: url,
    headers: {
        Authorization: Auth.authHeader()
    }
})

export default Service
