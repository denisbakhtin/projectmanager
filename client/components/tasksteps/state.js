import m from 'mithril'
import {
    addSuccess
} from '../shared/notifications'
import {
    responseErrors
} from '../../utils/helpers'
import Auth from '../../utils/auth'
import service from '../../utils/service.js'

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
        return service.getTaskStep(state.step.id)
            .then((result) => state.step = result)
            .catch((error) => state.errors = responseErrors(error))
    },
    getAll() {
        return service.getTaskSteps()
            .then((result) => state.steps = result.slice(0))
            .catch((error) => state.errors = responseErrors(error))
    },
    create() {
        return service.createTaskStep(state.step)
            .then((result) => {
                addSuccess("Task step created.")
                m.route.set('/task_steps')
            })
            .catch((error) => state.errors = responseErrors(error))
    },
    update() {
        return service.updateTaskStep(state.step.id, state.step)
            .then((result) => {
                addSuccess("Task step updated.")
                m.route.set('/task_steps')
            })
            .catch((error) => state.errors = responseErrors(error))
    },
    destroy() {
        return service.deleteTaskStep(state.step.id)
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
