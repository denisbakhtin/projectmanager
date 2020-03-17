import m from 'mithril'
import error from '../shared/error'
import {
    addSuccess
} from '../shared/notifications'
import {
    responseErrors
} from '../../utils/helpers'
import service from '../../utils/service'

export default function Role() {
    let role = {},
        errors = [],
        isNew = true,
        setName = (name) => role.name = name,

        //requests
        get = (id) =>
            service.getRole(id)
                .then((result) => role = result)
                .catch((error) => errors = responseErrors(error)),

        create = () =>
            service.createRole(role)
                .then((result) => {
                    addSuccess("User role created.")
                    m.route.set('/roles')
                })
                .catch((error) => errors = responseErrors(error)),

        update = () =>
            service.updateRole(role.id, role)
                .then((result) => {
                    addSuccess("User role updated.")
                    m.route.set('/roles')
                })
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            isNew = (m.route.param('id') == undefined)
            if (!isNew)
                get(m.route.param('id'))
        },

        view(vnode) {
            return m(".role", [
                m('h1.mb-4', (isNew) ? "Create user role" : 'Edit user role'),
                m('.form-group', [
                    m('label', 'Role name'),
                    m('input.form-control[type=text]', {
                        oncreate: (el) => el.dom.focus(),
                        oninput: (e) => setName(e.target.value),
                        value: role.name
                    })
                ]),
                m('.mb-2', m(error, {
                    errors: errors
                })),
                m('.actions', [
                    m('button.btn.btn-primary.mr-2[type=button]', {
                        onclick: (isNew) ? create : update
                    }, "Save"),
                    m('button.btn.btn-secondary[type=button]', {
                        onclick: () => window.history.back()
                    }, "Cancel")
                ]),
            ])
        }
    }
}
