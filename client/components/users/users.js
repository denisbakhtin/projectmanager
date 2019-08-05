import m from 'mithril'
import Auth from '../../utils/auth'
import { responseErrors } from '../../utils/helpers'
import error from '../shared/error'

const state = {
    users: [],
    errors: [],
    get() {
        state.errors = []
        m.request({
            method: "GET",
            url: "/api/users",
            headers: { Authorization: Auth.authHeader() }
        }).then((result) => {
            state.users = result.slice(0)
        }).catch((error) => state.errors = responseErrors(error))
    },
}

const Users = {
    oninit(vnode) {
        state.get()
    },
    
    view(vnode) {
        let ui = vnode.state;
        return m(".users", [
            m('h1.mb-4', 'Users'),
            m('table.table', [
                m('thead', [
                    m('tr', [
                        m('th[scope=col]', 'Name'),
                        m('th[scope=col]', 'Email'),
                        m('th.shrink.text-center[scope=col]', 'Actions')
                    ])
                ]),
                m('tbody', [
                    state.errors.length ? m('tr', m('td[colspan=3]', m(error, { errors: state.errors }))) : null,
                    state.users ? 
                        state.users.map((user) => {
                            return m('tr', { key: user.id }, [
                                m('td', user.name),
                                m('td', user.email),
                                m('td')
                            ])
                        }) : null
                ])
            ]),
        ])
    }
}

export default Users;