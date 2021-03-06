﻿import m from 'mithril'
import Auth from '../../utils/auth'
import {
    responseErrors
} from '../../utils/helpers'
import error from '../shared/error'
import service from '../../utils/service.js'

export default function Users() {
    let users = [],
        errors = [],

        getAll = () =>
            service.getUsers()
                .then((result) => {
                    users = result.slice(0)
                }).catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            getAll()
        },

        view(vnode) {
            return m(".users", [
                m('h1.title.mb-4', 'Users'),
                m('table.table', [
                    m('thead', [
                        m('tr', [
                            m('th[scope=col].hide-mobile', 'Name'),
                            m('th[scope=col]', 'Email'),
                            m('th.shrink.text-center[scope=col]', 'Actions')
                        ])
                    ]),
                    m('tbody', [
                        errors.length ? m('tr', m('td[colspan=3]', m(error, { errors: errors }))) : null,
                        users ? users.map((user) => {
                            return m('tr', {
                                key: user.id
                            }, [
                                m('td.hide-mobile', user.name),
                                m('td', user.email),
                                m('td')
                            ])
                        }) : null
                    ])
                ]),
            ])
        }
    }
}
