import m from 'mithril'
import error from '../shared/error'
import service from '../../utils/service'
import {
    responseErrors
} from '../../utils/helpers'

export default function Statuses() {
    let errors = [],
        statuses = [],

        getAll = () =>
        service.getStatuses()
        .then((result) => statuses = result.slice(0))
        .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            getAll()
        },

        view(vnode) {
            return m(".statuses", [
                m('h1.mb-4', 'Project status'),
                m('table.table', [
                    m('thead', [
                        m('tr', [
                            m('th[scope=col]', 'Name'),
                            m('th[scope=col]', 'Description'),
                            m('th[scope=col]', 'Order'),
                            m('th.shrink.text-center[scope=col]', 'Actions')
                        ])
                    ]),
                    m('tbody', [
                        statuses ?
                        statuses.map((status) => {
                            return m('tr', {
                                key: status.id
                            }, [
                                m('td', status.name),
                                m('td', status.description),
                                m('td', status.order),
                                m('td.shrink.text-center', m('button.btn.btn-outline-primary.btn-sm[type=button]', {
                                    onclick: () => {
                                        m.route.set('/statuses/edit/' + status.id)
                                    }
                                }, m('i.fa.fa-pencil')))
                            ])
                        }) : null
                    ])
                ]),
                errors.length ? m(error, {
                    errors: errors
                }) : null,
                m('.actions.mt-4', [
                    m('button.btn.btn-primary[type=button]', {
                        onclick: () => {
                            m.route.set('/statuses/new')
                        }
                    }, "New status")
                ]),
            ])
        }
    }
}
