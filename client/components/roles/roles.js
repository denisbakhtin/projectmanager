import m from 'mithril'
import Auth from '../../utils/auth'
import { responseErrors } from '../../utils/helpers'
import newRole from './new_role'
import editRole from './edit_role'
import destroyRole from './destroy_role'
import error from '../shared/error'

const state = {
    roles: [],
    errors: [],
    editRole: {},
    creating: false,
    destroyRole: {},
    get() {
        state.errors = []
        m.request({
            method: "GET",
            url: "/api/roles",
            headers: { Authorization: Auth.authHeader() }
        }).then((result) => {
            state.roles = result.slice(0)
        }).catch((error) => state.errors = responseErrors(error))
    },
}

const Roles = {
    oninit(vnode) {
        state.get()
    },
    
    view(vnode) {
        let ui = vnode.state;
        return m(".roles", [
            m('h1.mb-4', 'User roles'),
            m('table.table', [
                m('thead', [
                    m('tr', [
                        m('th[scope=col]', 'Name'),
                        m('th.shrink.text-center[scope=col]', 'Actions')
                    ])
                ]),
                m('tbody', [
                    state.errors.length ? m('tr', m('td[colspan=2]', m(error, { errors: state.errors }))) : null,
                    state.roles ? 
                        state.roles.map((role) => {
                            return m('tr', { key: role.id }, state.editRole.id == role.id ? [
                                m('td[colspan=2]', [
                                    m(editRole, { role: role, onUpdate: () => { state.editRole = {}; state.get(); }, onCancel: () => { state.editRole = {}}})
                                ])
                            ] : [
                                m('td', role.name),
                                m('td', m('.btn-group', [
                                    m('button.btn.btn-outline-primary.btn-sm[type=button]', { onclick: () => { state.editRole = role } }, m('i.fa.fa-pencil')),
                                    m('button.btn.btn-outline-danger.btn-sm[type=button]', { onclick: () => { state.destroyRole = role } }, m('i.fa.fa-times'))
                                ]))
                            ])
                        }) : null
                ])
            ]),
            //create new role
            state.creating ? m(newRole, { onCreate: () => { state.creating = false; state.get(); }, onCancel: () => { state.creating = false } }) : null,
            m('.actions', [
                m('button.btn.btn-primary[type=button]', { onclick: () => { state.creating = true } }, "New role")
            ]),
            //confirm role removal
            Object.keys(state.destroyRole).length ?
                m(destroyRole, { role: state.destroyRole, onDestroy: () => { state.destroyRole = {}; state.get(); }, onCancel: () => { state.destroyRole = {} } }) : null,
        ])
    }
}

export default Roles;