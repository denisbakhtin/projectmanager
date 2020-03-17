import m from 'mithril'
import error from '../shared/error'
import {
    responseErrors
} from '../../utils/helpers'
import service from '../../utils/service'

export default function Roles() {
    let roles = [],
        errors = [],

        getAll = () =>
            service.getRoles()
                .then((result) => roles = result.slice(0))
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            getAll()
        },

        view(vnode) {
            return m(".roles", [
                m('h1.mb-4', 'User roles'),
                m('table.table', [
                    m('thead', [
                        m('tr', [
                            m('th.shrink[scope=col]', '#'),
                            m('th[scope=col]', 'Name'),
                            m('th.shrink.text-center[scope=col]', 'Actions')
                        ])
                    ]),
                    m('tbody', [
                        roles ?
                            roles.map((role) => {
                                return m('tr', {
                                    key: role.id
                                }, [
                                    m('td', role.id),
                                    m('td', role.name),
                                    m('td.shrink.text-center', m('button.btn.btn-outline-primary.btn-sm[type=button]', {
                                        onclick: () => m.route.set('/roles/edit/' + role.id)
                                    },
                                        m('i.fa.fa-pencil'))
                                    )
                                ])
                            }) : null
                    ])
                ]),
                m(error, { errors: errors }),
                m('.actions.mt-4', [
                    m('button.btn.btn-primary[type=button]', {
                        onclick: () => m.route.set('/roles/new')
                    }, "New role")
                ]),
            ])
        }
    }
}
