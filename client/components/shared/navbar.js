import m from 'mithril'
import notifications from './navbar_notifications'
import profile from './navbar_profile'
import support from './navbar_support'
import curtime from './navbar_time'

const Navbar = {
    view(vnode) {
        return m('nav.navbar.navbar-expand-lg.navbar-light.bg-light.fixed-top', [
            m('a.navbar-brand[href=#!/]', [
                m('img', {
                    src: "/public/images/navbar_logo.png",
                    class: 'mr-2',
                    width: '23px',
                    height: '23px',
                }),
                "Project Manager",
            ]),
            m('button.sidebar-toggler.mr-4', m('span.fa.fa-bars')),
            m('form.form-inline.search-form.mr-2', [
                m('.input-group', [
                    m('input.form-control[type=search][placeholder=Search...]'),
                    m('.input-group-append', [
                        m('button.btn.btn-outline-primary[type=button]', m('i.fa.fa-search')),
                    ]),
                ]),
            ]),
            m('ul.navbar-nav.ml-auto', [
                m(curtime),
                m(notifications),
                m(profile),
                m(support),
            ])
        ])
    }
}

export default Navbar;