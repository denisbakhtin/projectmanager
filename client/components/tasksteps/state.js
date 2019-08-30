import m from 'mithril'
import {
    addSuccess
} from '../shared/notifications'
import {
    responseErrors
} from '../../utils/helpers'
import Auth from '../../utils/auth'

const state = {
    step: {},
    steps: [],
    errors: [],
    setName(name) {
        state.step.name = name
    },
    setIsfinal(is_final) {
        state.step.is_final = is_final
    },
    setOrder(order) {
        state.step.order = order
    },
    validate() {
        state.errors = []
        if (!state.step.name)
            state.errors.push("Step name is required.")
        return state.errors.length == 0
    },
    //requests
    get() {
        return m.request({
                method: "GET",
                url: "/api/task_steps/" + state.step.id,
                headers: {
                    Authorization: Auth.authHeader()
                }
            })
            .then((result) => state.step = result)
            .catch((error) => state.errors = responseErrors(error))
    },
    getAll() {
        return m.request({
                method: "GET",
                url: "/api/task_steps",
                headers: {
                    Authorization: Auth.authHeader()
                }
            })
            .then((result) => state.steps = result.slice(0))
            .catch((error) => state.errors = responseErrors(error))
    },
    create() {
        return m.request({
                method: "POST",
                url: "/api/task_steps",
                body: state.step,
                headers: {
                    Authorization: Auth.authHeader()
                }
            })
            .then((result) => {
                addSuccess("Task step created.")
                m.route.set('/task_steps')
            })
            .catch((error) => state.errors = responseErrors(error))
    },
    update() {
        return m.request({
                method: "PUT",
                url: "/api/task_steps/" + state.step.id,
                body: state.step,
                headers: {
                    Authorization: Auth.authHeader()
                }
            })
            .then((result) => {
                addSuccess("Task step updated.")
                m.route.set('/task_steps')
            })
            .catch((error) => state.errors = responseErrors(error))
    },
    destroy() {
        return m.request({
                method: "DELETE",
                url: "/api/task_steps/" + state.step.id,
                headers: {
                    Authorization: Auth.authHeader()
                }
            })
            .then((result) => {
                addSuccess("Task step removed.")
                m.route.set('/task_steps', {}, {
                    replace: true
                })
            })
            .catch((error) => state.errors = responseErrors(error))
    }
}

export default state