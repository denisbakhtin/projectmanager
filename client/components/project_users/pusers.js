import m from 'mithril'
import {
    ISODateToHtml5
} from '../../utils/helpers'
import error from '../shared/error'
import state from './state'

const ProjectUsers = {
    oninit(vnode) {
        state.pusers = []
        state.errors = []
        state.pickup = false
        state.getUsers()
        if (vnode.attrs.project_users && vnode.attrs.project_users.length > 0) state.pusers = vnode.attrs.project_users.slice(0)
        state.getRoles()
    },
    onchange(vnode) {
        if (typeof vnode.attrs.onchange == 'function') vnode.attrs.onchange(state.pusers)
    },

    view(vnode) {
        let ui = vnode.state
        let availableUsers = state.users.filter(u => state.pusers.findIndex(pu => pu.user_id == u.id) == -1)
        return m(".project_users", [
            state.pickup ? m('.pickup-form', [
                m('.row.align-items-center.mb-2', [
                    m('.col', [
                        m('h6.text-center', 'Project users'),
                        m('.card', [
                            m('.card-body', [
                                state.pusers.length > 0 ? m('table.no-borders.w-100', [
                                    state.pusers.map((puser) => {
                                        return m('tr', [
                                            m('td.user_entry', {
                                                onclick: () => {
                                                    state.removeUserFromProject(puser);
                                                    ui.onchange(vnode)
                                                }
                                            }, `${puser.user.name} (${puser.user.email})`),
                                            m('td.shrink.role-select', [
                                                state.roles.length > 0 ?
                                                m('select.form-control', {
                                                    onchange: function (e) {
                                                        val = e.target.value;
                                                        puser.role_id = val;
                                                        puser.role = state.roles.find((el) => {
                                                            return el.id == val
                                                        })
                                                    },
                                                    value: puser.role_id
                                                }, state.roles.map((role) => {
                                                    return m('option', {
                                                        value: role.id
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
                                            state.addUserToProject(vnode.attrs.project_id, user);
                                            ui.onchange(vnode)
                                        }
                                    }, `${user.name} (${user.email})`)
                                }) : m('.text-muted.text-center', "empty")
                            ])
                        ])

                    ])
                ]),
                m('.mb-2', m(error, {
                    errors: state.errors
                })),
                m('a.btn.btn-sm.btn-secondary[href=#]', {
                    onclick: () => {
                        state.pickup = false;
                        return false;
                    }
                }, 'Done')
            ]) : [
                m('a[href=#]', {
                    onclick: () => {
                        state.pickup = true;
                        return false;
                    }
                }, state.pusers.length > 0 ? state.pusers.map((u) => {
                    return `${u.user.name} (${u.role.name})`
                }).join(', ') : "None"),
                m('i.fa.fa-pencil.ml-2')
            ]

        ])
    }
}

export default ProjectUsers;