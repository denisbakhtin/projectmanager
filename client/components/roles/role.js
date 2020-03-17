import m from 'mithril'
import {
    responseErrors
} from '../../utils/helpers'
import {
    addSuccess
} from '../shared/notifications'

export default function Role() {
    let role = {},
        errors = [],

        //requests
        get = (id) =>
            service.getRole(id)
                .then((result) => role = result)
                .catch((error) => errors = responseErrors(error)),

        remove = () =>
            service.deleteRole(role.id)
                .then((result) => {
                    addSuccess("User role removed.")
                    m.route.set('/roles', {}, {
                        replace: true
                    })
                })
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            get(m.route.param('id'))
        },

        view(vnode) {
            return m(".role", role ? [
                m('h1.mb-2', role.name),
                m('.actions', [
                    m('button.btn.btn-primary.mr-2[type=button]', {
                        onclick: () => m.route.set('/roles/edit/' + role.id)
                    }, "Edit"),
                    m('button.btn.btn-secondary.mr-2[type=button]', {
                        onclick: () => m.route.set('/roles')
                    }, "Back to list"),
                    m('button.btn.btn-outline-danger[type=button]', {
                        onclick: remove
                    }, "Remove step")
                ])
            ] : m('Loading...'))
        }
    }
}
