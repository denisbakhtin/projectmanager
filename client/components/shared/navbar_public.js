import m from 'mithril';

const Navbar = {
    view(vnode) {
        return m('nav.navbar.navbar-expand-lg.navbar-dark.navbar-bg', [
            m('.container', [
                m('a.navbar-brand[href=#!/]', [
                    m('img', {
                        src: "/public/images/navbar_logo.png",
                        class: 'mr-2',
                        width: '23px',
                        height: '23px',
                    }),
                    "Project Manager",
                ]),
                m('button.navbar-toggler[type=button][data-toggle=collapse][data-target=#navbarContent]', m('span.navbar-toggler-icon')),
                m('#navbarContent.collapse.navbar-collapse', [
                    m('ul.navbar-nav ml-auto', [
                        m('li.nav-item', m('a.nav-link[href=#!/login]', 'Login')),
                        m('li.nav-item', m('a.nav-link[href=#!/register]', 'Register')),
                    ])
                ])
            ])
        ])
    }
}
export default Navbar;