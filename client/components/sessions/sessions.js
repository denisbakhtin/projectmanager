import m from 'mithril'
import {
    responseErrors
} from '../../utils/helpers'
import error from '../shared/error'
import service from '../../utils/service.js'
import sessions_item from './sessions_item'

export default function Sessions() {
    let sessions = [],
        errors = [],

        getAll = () =>
            service.getSessions()
                .then((result) => sessions = result.slice(0))
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            getAll()
        },

        view(vnode) {
            return m(".sessions", [
                m('h1.title.mb-4', 'Sessions'),
                sessions.length > 0 ?
                    m('ul.dashboard-box.box-list', [
                        sessions.map((session) => m(sessions_item, { key: session.id, session: session, onUpdate: getAll }))
                    ]) : m('p', 'No sessions yet.'),
                m(error, { errors: errors }),
                m('button#floating-add-button.btn.btn-primary[type=button]', {
                    onclick: () => m.route.set('/sessions/new')
                },
                    m('i.fa.fa-plus'),
                ),
            ])
        }
    }
}
