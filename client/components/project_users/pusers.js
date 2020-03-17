import m from 'mithril'
import {
    ISODateToHtml5
} from '../../utils/helpers'
import error from '../shared/error'
import {
    addSuccess
} from '../shared/notifications'
import {
    responseErrors
} from '../../utils/helpers'
import Auth from '../../utils/auth'
import service from '../../utils/service.js'

export default function ProjectUsers() {
    let users = [],
        pusers = [],
        roles = [],
        errors = [],
        pickup = false,

        //requests
        getUsers = () =>
            service.getUsers()
                .then((result) => {
                    users = result.slice(0)
                    let authuser = Auth.getAuthenticatedUser
                    let index = users.findIndex(x => x.id === authuser.user_id)
                    users.splice(index, 1)
                })
                .catch((error) => errors = responseErrors(error)),

        addUserToProject = (project_id, user) => {
            let role = roles.length > 0 ? roles[0] : {}
            pusers.push({
                user_id: user.id,
                role_id: role.id,
                project_id,
                user,
                role
            })
        },

        removeUserFromProject = (puser) => {
            let index = pusers.findIndex(x => x.user_id === puser.user_id)
            pusers.splice(index, 1)
        },

        getRoles = () =>
            service.getRoles()
                .then((result) => roles = result.slice(0))
                .catch((error) => errors = responseErrors(error)),

        onchange = (vnode) => {
            if (typeof vnode.attrs.onchange == 'function') vnode.attrs.onchange(pusers)
        }

    return {
        oninit(vnode) {
            getUsers()
            if (vnode.attrs.project_users && vnode.attrs.project_users.length > 0) pusers = vnode.attrs.project_users.slice(0)
            getRoles()
        },

        view(vnode) {
            let availableUsers = users.filter(u => pusers.findIndex(pu => pu.user_id == u.id) == -1)
            return m(".project_users", [
                pickup ? m('.pickup-form', [
                    m('.row.align-items-center.mb-2', [
                        m('.col', [
                            m('h6.text-center', 'Project users'),
                            m('.card', [
                                m('.card-body', [
                                    pusers.length > 0 ? m('table.no-borders.w-100', [
                                        pusers.map((puser) => {
                                            return m('tr', [
                                                m('td.user_entry', {
                                                    onclick: () => {
                                                        removeUserFromProject(puser);
                                                        onchange(vnode)
                                                    }
                                                }, `${puser.user.name} (${puser.user.email})`),
                                                m('td.shrink.role-select', [
                                                    roles.length > 0 ?
                                                        m('select.form-control', {
                                                            onchange: function (e) {
                                                                let val = e.target.value;
                                                                puser.role_id = val;
                                                                puser.role = roles.find((el) => (el.id == val));
                                                            },
                                                            value: puser.role_id
                                                        }, roles.map((role) => {
                                                            return m('option', {
                                                                value: role.id,
                                                                selected: (role.id == puser.role_id)
                                                            }, role.name)
                                                        })) : null
                                                ])
                                            ])
                                        })
                                    ]) : m('.text-muted.text-center', "empty")
                                ])
                            ])
                        ]),
                        m('.col-sm-1', [
                            m('.text-center.mt-4', m('i.fa.fa-2x.fa-arrows-h')),
                        ]),
                        m('.col', [
                            m('h6.text-center', 'Available users'),
                            m('.card', [
                                m('.card-body', [
                                    availableUsers.length > 0 ? availableUsers.map((user) => {
                                        return m('.user_entry', {
                                            onclick: () => {
                                                addUserToProject(vnode.attrs.project_id, user);
                                                onchange(vnode)
                                            }
                                        }, `${user.name} (${user.email})`)
                                    }) : m('.text-muted.text-center', "empty")
                                ])
                            ])
                        ])
                    ]),
                    m('.mb-2', m(error, {
                        errors: errors
                    })),
                    m('a.btn.btn-sm.btn-secondary[href=#]', {
                        onclick: () => {
                            pickup = false;
                            return false;
                        }
                    }, 'Done')
                ]) : [
                        m('a[href=#]', {
                            onclick: () => {
                                pickup = true;
                                return false;
                            }
                        }, pusers.length > 0 ? pusers.map((u) => {
                            return `${u.user.name} (${u.role.name})`
                        }).join(', ') : "None"),
                        m('i.fa.fa-pencil.ml-2')
                    ]
            ])
        }
    }
}
