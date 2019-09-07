import m from 'mithril'
import {
    responseErrors
} from '../../utils/helpers'
import service from '../../utils/service.js'
import {
    addSuccess
} from '../shared/notifications'

export default function Status() {
    var status = {},
        errors = []

    function get() {
        return service.getStatus(status.id)
            .then((result) => status = result)
            .catch((error) => errors = responseErrors(error))
    }
    function destroy() {
        return service.deleteStatus(status.id)
            .then((result) => addSuccess("Status removed."))
            .catch((error) => errors = responseErrors(error))
    }
    return {
        oninit(vnode) {
            errors = []
            status = { id: m.route.param('id') }
            get()
        },

        view(vnode) {
            return m(".status", [
                status.name ? [
                    m('h1.mb-2', status.name),
                    status.description ? [
                        m('h3', "Description"),
                        m('p', status.description)
                    ] : null
                ] : null, 
                m('.actions', [
                    m('button.btn.btn-primary.mr-2[type=button]', { onclick: () => { m.route.set('/statuses/edit/' + status.id) } }, "Edit"),
                    m('button.btn.btn-secondary.mr-2[type=button]', { onclick: () => { m.route.set('/statuses') } }, "Back to list"),
                    m('button.btn.btn-outline-danger[type=button]', { onclick: destroy }, "Remove status")
                ])
            ])
        }
    }
}
