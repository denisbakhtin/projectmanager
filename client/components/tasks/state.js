import m from 'mithril'
import {
    addSuccess
} from '../shared/notifications'
import {
    responseErrors
} from '../../utils/helpers'
import Auth from '../../utils/auth'

const state = {
    task: {},
    tasks: [],
    errors: [],
    setName(name) {
        state.task.name = name
    },
    setDescription(description) {
        state.task.description = description
    },
    setStepID(step_id) {
        state.task.task_step_id = step_id
    },
    setProjectUserID(project_user_id) {
        state.task.project_user_id = project_user_id
    },
    setProjectID(project_id) {
        state.task.project_id = project_id
    },
    validate() {
        state.errors = []
        if (!state.task.name)
            state.errors.push("Task name is required.")
        return state.errors.length == 0
    },
    //requests
    get() {
        return m.request({
                method: "GET",
                url: "/api/tasks/" + state.task.id,
                headers: {
                    Authorization: Auth.authHeader()
                }
            })
            .then((result) => state.task = result)
            .catch((error) => state.errors = responseErrors(error))
    },
    getAll() {
        return m.request({
                method: "GET",
                url: "/api/tasks",
                headers: {
                    Authorization: Auth.authHeader()
                }
            })
            .then((result) => state.tasks = result.slice(0))
            .catch((error) => state.errors = responseErrors(error))
    },
    create() {
        return m.request({
                method: "POST",
                url: "/api/tasks",
                body: state.task,
                headers: {
                    Authorization: Auth.authHeader()
                }
            })
            .then((result) => {
                addSuccess("Task created.")
                m.route.set('/tasks')
            })
            .catch((error) => state.errors = responseErrors(error))
    },
    update() {
        return m.request({
                method: "PUT",
                url: "/api/tasks/" + state.task.id,
                body: state.task,
                headers: {
                    Authorization: Auth.authHeader()
                }
            })
            .then((result) => {
                addSuccess("Task updated.")
                m.route.set('/tasks')
            })
            .catch((error) => state.errors = responseErrors(error))
    },
    destroy() {
        return m.request({
                method: "DELETE",
                url: "/api/tasks/" + state.task.id,
                headers: {
                    Authorization: Auth.authHeader()
                }
            })
            .then((result) => {
                addSuccess("Task removed.")
                m.route.set('/tasks', {}, {
                    replace: true
                })
            })
            .catch((error) => state.errors = responseErrors(error))
    }
}

export default state