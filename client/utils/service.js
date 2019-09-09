import m from 'mithril'
import {
    responseErrors,
    ISODateToHtml5
} from './helpers'
import Auth from './auth'

const Service = {
    //>>>>>>>>>>>>> roles
    getRoles: () =>
        m.request({
            method: "GET",
            url: "/api/roles",
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    createRole: (name) =>
        m.request({
            method: "POST",
            url: "/api/roles",
            body: {
                name: name
            },
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    updateRole: (id, name) =>
        m.request({
            method: "PUT",
            url: "/api/roles/" + id,
            body: {
                id: id,
                name: name
            },
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    deleteRole: (id) =>
        m.request({
            method: "DELETE",
            url: "/api/roles/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    //>>>>>>>>>>>>> projects
    getProject: (id) =>
        m.request({
            method: "GET",
            url: "/api/projects/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    getProjects: () =>
        m.request({
            method: "GET",
            url: "/api/projects",
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    newProject: () =>
        m.request({
            method: "GET",
            url: "/api/new_project/",
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    createProject: (project) =>
        m.request({
            method: "POST",
            url: "/api/projects",
            body: project,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    editProject: (id) =>
        m.request({
            method: "GET",
            url: "/api/edit_project/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    updateProject: (id, project) =>
        m.request({
            method: "PUT",
            url: "/api/projects/" + id,
            body: project,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    deleteProject: (id) =>
        m.request({
            method: "DELETE",
            url: "/api/projects/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),

    //>>>>>>>>>>>>>>>>>> tasks
    getTask: (id) =>
        m.request({
            method: "GET",
            url: "/api/tasks/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    getTasks: () =>
        m.request({
            method: "GET",
            url: "/api/tasks",
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    newTask: () =>
        m.request({
            method: "GET",
            url: "/api/new_task/",
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    createTask: (task) =>
        m.request({
            method: "POST",
            url: "/api/tasks",
            body: task,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    editTask: (id) =>
        m.request({
            method: "GET",
            url: "/api/edit_task/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    updateTask: (id, task) =>
        m.request({
            method: "PUT",
            url: "/api/tasks/" + id,
            body: task,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    deleteTask: (id) =>
        m.request({
            method: "DELETE",
            url: "/api/tasks/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),

    //>>>>>>>>>>>>>> statuses
    getStatus: (id) =>
        m.request({
            method: "GET",
            url: "/api/statuses/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    getStatuses: () =>
        m.request({
            method: "GET",
            url: "/api/statuses",
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    createStatus: (status) =>
        m.request({
            method: "POST",
            url: "/api/statuses",
            body: status,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    updateStatus: (id, status) =>
        m.request({
            method: "PUT",
            url: "/api/statuses/" + id,
            body: status,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    deleteStatus: (id) =>
        m.request({
            method: "DELETE",
            url: "/api/statuses/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),

    //>>>>>>>>>>>> notifications
    getNotifications: () =>
        m.request({
            method: "GET",
            url: "/api/notifications",
            headers: {
                Authorization: Auth.authHeader()
            },
        }),
    deleteNotification: (id) =>
        m.request({
            method: "DELETE",
            url: "/api/notifications/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),

    //>>>>>>>>>>>>>>> task steps
    getTaskStep: (id) =>
        m.request({
            method: "GET",
            url: "/api/task_steps/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    getTaskSteps: () =>
        m.request({
            method: "GET",
            url: "/api/task_steps",
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    createTaskStep: (step) =>
        m.request({
            method: "POST",
            url: "/api/task_steps",
            body: step,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    updateTaskStep: (id, step) =>
        m.request({
            method: "PUT",
            url: "/api/task_steps/" + id,
            body: step,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    deleteTaskStep: (id) =>
        m.request({
            method: "DELETE",
            url: "/api/task_steps/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),

    //>>>>>>>>>>>>> users
    getUsers: () =>
        m.request({
            method: "GET",
            url: "/api/users",
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    getProjectUsers: (project_id) =>
        m.request({
            method: "GET",
            url: "/api/project_users/" + project_id,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    getRoles: () =>
        m.request({
            method: "GET",
            url: "/api/roles",
            headers: {
                Authorization: Auth.authHeader()
            }
        }),

    //>>>>>>>>>>>>>> settings
    getSetting: (id) =>
        m.request({
            method: "GET",
            url: "/api/settings/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    getSettings: () =>
        m.request({
            method: "GET",
            url: "/api/settings",
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    createSetting: (setting) =>
        m.request({
            method: "POST",
            url: "/api/settings",
            body: setting,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    updateSetting: (id, setting) =>
        m.request({
            method: "PUT",
            url: "/api/settings/" + id,
            body: setting,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    deleteSetting: (id) =>
        m.request({
            method: "DELETE",
            url: "/api/settings/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        }),
    //>>>>>>>>>>>> upload file
    uploadFile: (file) => {
        let data = new FormData()
        data.append("upload", file)
        return m.request({
            method: "POST",
            url: `/api/upload/form`,
            body: data,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },

}

export default Service
