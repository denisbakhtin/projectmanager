import m from 'mithril'
import error from '../shared/error'
import {
    addSuccess
} from '../shared/notifications'
import {
    responseErrors
} from '../../utils/helpers'
import service from '../../utils/service.js'

export default function Status() {
    let errors = [],
        status = {},
        isNew = true,

        setName = (name) => status.name = name,
        setDescription = (description) => status.description = description,
        setOrd = (order) => status.ord = order,

        validate = () => {
            errors = []
            if (!status.name)
                errors.push("Status name is required.")
            return errors.length == 0
        },

        get = () =>
        service.getStatus(status.id)
        .then((result) => status = result)
        .catch((error) => errors = responseErrors(error)),

        create = () =>
        service.createStatus(status)
        .then((result) => {
            addSuccess("Status created.")
            m.route.set('/statuses')
        })
        .catch((error) => errors = responseErrors(error)),

        update = () =>
        service.updateStatus(status.id, status)
        .then((result) => {
            addSuccess("Status updated.")
            m.route.set('/statuses')
        })
        .catch((error) => errors = responseErrors(error))

        return {
            oninit(vnode) {
                if (m.route.param('id')) {
                    isNew = false
                    status = {
                        id: m.route.param('id')
                    }
                    get()
                } else
                    status = {
                        ord: "1"
                    }
                errors = []
            },

            view(vnode) {
                return m(".statuses", [
                    m('h1.mb-4', (isNew) ? 'New status' : 'Edit status'),
                    m('.form-group', [
                        m('label', 'Status name'),
                        m('input.form-control[type=text]', {
                            oncreate: (el) => {
                                el.dom.focus()
                            },
                            oninput: (e) => setName(e.target.value),
                            value: status.name
                        })
                    ]),
                    m('.form-group w-25', [
                        m('label', 'Order'),
                        m('input.form-control[type=number][min=0]', {
                            oninput: (e) => setOrd(e.target.value),
                            value: status.ord
                        })
                    ]),
                    m('.form-group', [
                        m('label', "Description"),
                        m('textarea.form-control', {
                            oninput: (e) => setDescription(e.target.value),
                            value: status.description
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
                            onclick: () => {
                                window.history.back()
                            }
                        }, "Cancel")
                    ]),
                ])
            }
        }
}
