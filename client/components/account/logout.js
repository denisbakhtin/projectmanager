import m from 'mithril'
import Auth from '../../utils/auth'

export default function Logout() {
    return {
        oninit(vnode) {
            Auth.logout()
        },
        view(vnode) {
            return null
        }
    }
}
