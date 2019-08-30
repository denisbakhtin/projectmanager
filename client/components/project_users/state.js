import m from 'mithril'
import { addSuccess } from '../shared/notifications'
import { responseErrors } from '../../utils/helpers'
import Auth from '../../utils/auth'

const state = {
    users: [],
    pusers: [],
    roles: [],
    errors: [],
    pickup: false,
    //requests
    getUsers() {
        m.request({
            method: "GET",
            url: "/api/users",
            headers: { Authorization: Auth.authHeader() }
        })
            .then((result) => {
                state.users = result.slice(0)
                let authuser = Auth.getAuthenticatedUser
                let index = state.users.findIndex(x => x.id === authuser.user_id)
                state.users.splice(index, 1)
            })
            .catch((error) => state.errors = responseErrors(error))
    },
    addUserToProject(project_id, user) {
        let role = state.roles.length > 0 ? state.roles[0] : {}
        state.pusers.push({
            user_id: user.id, 
            role_id: role.id,
            project_id, 
            user,
            role
        })
    },
    removeUserFromProject(puser) {
        let index = state.pusers.findIndex(x => x.user_id === puser.user_id)
        state.pusers.splice(index, 1)
    },
    getRoles() {
        return m.request({
            method: "GET",
            url: "/api/roles",
            headers: { Authorization: Auth.authHeader() }
        })
            .then((result) => state.roles = result.slice(0))
            .catch((error) => state.errors = responseErrors(error))
    },
}

export default state