
import m from 'mithril'
import {
    responseErrors,
    ISODateToHtml5
} from './helpers'
import Auth from './auth'

const Service = {
    //>>>>>>>>>>>>> projects
    getProject(id) {
        return m.request({
            method: "GET",
            url: "/api/projects/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },
    getProjects() {
        return m.request({
            method: "GET",
            url: "/api/projects",
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },
    createProject(project){
        return m.request({
            method: "POST",
            url: "/api/projects",
            body: project,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },
    updateProject(id, project) {
        return m.request({
            method: "PUT",
            url: "/api/projects/" + id,
            body: project,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },
    deleteProject(id){
        m.request({
            method: "DELETE",
            url: "/api/projects/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },

    //>>>>>>>>>>>>>>>>>> tasks
    getTask(id) {
        return m.request({
            method: "GET",
            url: "/api/tasks/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },
    getTasks() {
        return m.request({
            method: "GET",
            url: "/api/tasks",
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },
    createTask(task){
        return m.request({
            method: "POST",
            url: "/api/tasks",
            body: task,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },
    updateTask(id, task){
        return m.request({
            method: "PUT",
            url: "/api/tasks/" + id,
            body: task,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },
    deleteTask(id) {
        return m.request({
            method: "DELETE",
            url: "/api/tasks/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },

    //>>>>>>>>>>>>>> statuses
    getStatus(id){
        return m.request({
            method: "GET",
            url: "/api/statuses/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },
    getStatuses() {
        return m.request({
            method: "GET",
            url: "/api/statuses",
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },
    createStatus(status){
        return m.request({
            method: "POST",
            url: "/api/statuses",
            body: status,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },
    updateStatus(id, status){
        return m.request({
            method: "PUT",
            url: "/api/statuses/" + id,
            body: status,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },
    deleteStatus(id){
        return m.request({
            method: "DELETE",
            url: "/api/statuses/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },

    //>>>>>>>>>>>> notifications
    getNotifications() {
        return m.request({
            method: "GET",
            url: "/api/notifications",
            headers: {
                Authorization: Auth.authHeader()
            },
        })
    },
    deleteNotification(id){
        return m.request({
            method: "DELETE",
            url: "/api/notifications/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },

    //>>>>>>>>>>>>>>> task steps
    getTaskStep(id){
        return m.request({
            method: "GET",
            url: "/api/task_steps/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },
    getTaskSteps() {
        return m.request({
            method: "GET",
            url: "/api/task_steps",
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },
    createTaskStep(step){
        return m.request({
            method: "POST",
            url: "/api/task_steps",
            body: step,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },
    updateTaskStep(id, step){
        return m.request({
            method: "PUT",
            url: "/api/task_steps/" + id,
            body: step,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },
    deleteTaskStep(id){
        return m.request({
            method: "DELETE",
            url: "/api/task_steps/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },

    //>>>>>>>>>>>>> users
    getUsers() {
        return m.request({
            method: "GET",
            url: "/api/users",
            headers: { Authorization: Auth.authHeader() }
        })
    },
    getProjectUsers(project_id){
        return m.request({
            method: "GET",
            url: "/api/project_users/"+project_id,
            headers: { Authorization: Auth.authHeader() }
        })
    },
    getRoles() {
        return m.request({
            method: "GET",
            url: "/api/roles",
            headers: { Authorization: Auth.authHeader() }
        })
    },

    //>>>>>>>>>>>>>> settings
    getSetting(id){
        return m.request({
            method: "GET",
            url: "/api/settings/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },
    getSettings() {
        return m.request({
            method: "GET",
            url: "/api/settings",
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },
    createSetting(setting){
        return m.request({
            method: "POST",
            url: "/api/settings",
            body: setting,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },
    updateSetting(id, setting){
        return m.request({
            method: "PUT",
            url: "/api/settings/" + id,
            body: setting,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },
    deleteSetting(id){
        return m.request({
            method: "DELETE",
            url: "/api/settings/" + id,
            headers: {
                Authorization: Auth.authHeader()
            }
        })
    },
}

export default Service
