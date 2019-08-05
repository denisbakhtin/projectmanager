import m from 'mithril'
import {
    addSuccess
} from '../shared/notifications'
import {
    responseErrors
} from '../../utils/helpers'
import Auth from '../../utils/auth'

const state = {
    status: {},
    statuses: [],
    errors: [],
    setName(name) {
        state.status.name = name
    },
    setDescription(description) {
        state.status.description = description
    },
    setOrder(order) {
        state.status.order = order
    },
    validate() {
        state.errors = []
        if (!state.status.name)
            state.errors.push("Status name is required.")
        return state.errors.length == 0
    },
    //requests
    get() {
        return m.request({
                method: "GET",
                url: "/api/statuses/" + state.status.id,
                headers: {
                    Authorization: Auth.authHeader()
                }
            })
            .then((result) => state.status = result)
            .catch((error) => state.errors = responseErrors(error))
    },
    getAll() {
        return m.request({
                method: "GET",
                url: "/api/statuses",
                headers: {
                    Authorization: Auth.authHeader()
                }
            })
            .then((result) => state.statuses = result.slice(0))
            .catch((error) => state.errors = responseErrors(error))
    },
    create() {
        return m.request({
                method: "POST",
                url: "/api/statuses",
                body: state.status,
                headers: {
                    Authorization: Auth.authHeader()
                }
            })
            .then((result) => {
                addSuccess("Status created.")
                m.route.set('/statuses')
            })
            .catch((error) => state.errors = responseErrors(error))
    },
    update() {
        return m.request({
                method: "PUT",
                url: "/api/statuses/" + state.status.id,
                body: state.status,
                headers: {
                    Authorization: Auth.authHeader()
                }
            })
            .then((result) => {
                addSuccess("Status updated.")
                m.route.set('/statuses')
            })
            .catch((error) => state.errors = responseErrors(error))
    },
    destroy() {
        return m.request({
                method: "DELETE",
                url: "/api/statuses/" + state.status.id,
                headers: {
                    Authorization: Auth.authHeader()
                }
            })
            .then((result) => {
                addSuccess("Status removed.")
                m.route.set('/statuses', {}, {
                    replace: true
                })
            })
            .catch((error) => state.errors = responseErrors(error))
    }
}

export default state