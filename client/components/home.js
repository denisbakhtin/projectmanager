import m from 'mithril'
import Auth from '../utils/auth'
import login from './account/login';

export default function Home() {
    return {
        view(vnode) {
            return Auth.isLoggedIn() ? [
                m('.home', [
                    m('h1', 'My Dashboard'),
                    m('p', `You are currently logged in as ${Auth.getAuthenticatedUser().name}.`),
                ])
            ] : m(login);
        }
    }
}
