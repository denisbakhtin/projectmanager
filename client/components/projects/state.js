import m from 'mithril'
import {
    addSuccess
} from '../shared/notifications'
import {
    responseErrors,
    ISODateToHtml5
} from '../../utils/helpers'
import Auth from '../../utils/auth'

const state = {
    project: {},
    projects: [],
    errors: [],
    setName(name) {
        state.project.name = name
    },
    setDescription(description) {
        state.project.description = description
    },
    setStatusId(status_id) {
        state.project.status_id = status_id
    },
    setStartDate(date) {
        state.project.start_date = date
        if (date && state.project.end_date && state.project.end_date < date)
            state.project.end_date = null
    },
    setEndDate(date) {
        state.project.end_date = date
        if (date && state.project.start_date && state.project.start_date > date)
            state.project.start_date = null
    },
    setProjectUsers(pusers) {
        state.project.project_users = pusers
    },
    setFiles(files) {
        state.project.files = files
    },
    validate() {
        state.errors = []
        if (!state.project.name)
            state.errors.push("Project name is required.")
        if (state.project.start_date && state.project.end_date && state.project.start_date > state.project.end_date)
            state.errors.push("End date cannot be earlier than start date.")
        return state.errors.length == 0
    },
    toGo(proj) {
        let obj = Object.assign(proj, {
            status_id: proj.status_id ? "" + proj.status_id : undefined,
            start_date: proj.start_date ? new Date(proj.start_date).toISOString() : undefined,
            end_date: proj.end_date ? new Date(proj.end_date).toISOString() : undefined,
            project_users: proj.project_users.map((pu) => {
                pu.role_id = "" + pu.role_id;
                return pu
            }),
            files: proj.files
        })
        return obj
    },
    fromGo(proj) {
        return Object.assign(proj, {
            start_date: ISODateToHtml5(proj.start_date || null),
            end_date: ISODateToHtml5(proj.end_date || null)
        })
    },
    getOwnerName() {
        return (state.project.owner || Auth.getAuthenticatedUser()).name
    },
    //requests
    get() {
        return m.request({
                method: "GET",
                url: "/api/projects/" + state.project.id,
                headers: {
                    Authorization: Auth.authHeader()
                }
            })
            .then((result) => state.project = state.fromGo(result))
            .catch((error) => state.errors = responseErrors(error))
    },
    getAll() {
        return m.request({
                method: "GET",
                url: "/api/projects",
                headers: {
                    Authorization: Auth.authHeader()
                }
            })
            .then((result) => state.projects = result.map((r) => state.fromGo(r)))
            .catch((error) => state.errors = responseErrors(error))
    },
    create() {
        return m.request({
                method: "POST",
                url: "/api/projects",
                body: state.toGo(state.project),
                headers: {
                    Authorization: Auth.authHeader()
                }
            })
            .then((result) => {
                addSuccess("Project created.")
                m.route.set('/projects')
            })
            .catch((error) => state.errors = responseErrors(error))
    },
    update() {
        return m.request({
                method: "PUT",
                url: "/api/projects/" + state.project.id,
                body: state.toGo(state.project),
                headers: {
                    Authorization: Auth.authHeader()
                }
            })
            .then((result) => {
                addSuccess("Project updated.")
                m.route.set('/projects')
            })
            .catch((error) => state.errors = responseErrors(error))
    },
    destroy() {
        return m.request({
                method: "DELETE",
                url: "/api/projects/" + state.project.id,
                headers: {
                    Authorization: Auth.authHeader()
                }
            })
            .then((result) => {
                addSuccess("Project removed.")
                m.route.set('/projects', {}, {
                    replace: true
                })
            })
            .catch((error) => state.errors = responseErrors(error))
    }
}

export default state