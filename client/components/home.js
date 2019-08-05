import m from 'mithril'
import Auth from '../utils/auth'
import login from './account/login';

const Home = {
    view(vnode) {
        return Auth.isLoggedIn() ? [
            m('.home', [
                m('h1', 'My Dashboard'),
                m('p', `You are currently logged in as ${Auth.getAuthenticatedUser().name}.`),
            ])
        ] : m(login);
    }
}

export default Home