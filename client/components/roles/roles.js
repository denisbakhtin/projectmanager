import m from 'mithril'
import Auth from '../../utils/auth'
import {
    responseErrors
} from '../../utils/helpers'
import newRole from './new_role'
import editRole from './edit_role'
import destroyRole from './destroy_role'
import error from '../shared/error'
import service from '../../utils/service.js'

export default function Roles() {
    let roles = [],
        errors = [],
        editRole = {},
        creating = false,
        destroyRole = {},

        get = () =>
        service.getRoles().then((result) => {
            roles = result.slice(0)
        }).catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            get()
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
                        errors.length ? m('tr', m('td[colspan=2]', m(error, {
                            errors: errors
                        }))) : null,
                        roles ?
                        roles.map((role) => {
                            return m('tr', {
                                key: role.id
                            }, editRole.id == role.id ? [
                                m('td[colspan=3]', [
                                    m(editRole, {
                                        role: role,
                                        onUpdate: () => {
                                            editRole = {};
                                            get();
                                        },
                                        onCancel: () => {
                                            editRole = {}
                                        }
                                    })
                                ])
                            ] : [
                                m('td.shrink', role.id),
                                m('td', role.name),
                                m('td', m('.btn-group', [
                                    m('button.btn.btn-outline-primary.btn-sm[type=button]', {
                                        onclick: () => {
                                            editRole = role
                                        }
                                    }, m('i.fa.fa-pencil')),
                                    m('button.btn.btn-outline-danger.btn-sm[type=button]', {
                                        onclick: () => {
                                            destroyRole = role
                                        }
                                    }, m('i.fa.fa-times'))
                                ]))
                            ])
                        }) : null
                    ])
                ]),
                //create new role
                creating ? m(newRole, {
                    onCreate: () => {
                        creating = false;
                        get();
                    },
                    onCancel: () => {
                        creating = false
                    }
                }) : null,
                m('.actions', [
                    m('button.btn.btn-primary[type=button]', {
                        onclick: () => {
                            creating = true
                        }
                    }, "New role")
                ]),
                //confirm role removal
                Object.keys(destroyRole).length ?
                m(destroyRole, {
                    role: destroyRole,
                    onDestroy: () => {
                        destroyRole = {};
                        get();
                    },
                    onCancel: () => {
                        destroyRole = {}
                    }
                }) : null,
            ])
        }
    }
}
