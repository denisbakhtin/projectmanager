﻿import m from 'mithril'
import notifications from './navbar_notifications'
import profile from './navbar_profile'
import curtime from './navbar_time'
import state from './state'

export default function Navbar() {
    return {
        view(vnode) {
            return m('nav.navbar.navbar-expand-md.navbar-light.bg-light.fixed-top', [
                m('a.navbar-brand[href=#!/]', [
                    m('img', {
                        src: "/public/images/navbar_logo.png",
                        class: 'mr-2',
                        width: '23px',
                        height: '23px',
                    }),
                    "Project Manager",
                ]),
                m('button.sidebar-toggler.mr-4', {
                    onclick: state.toggleSidebar
                }, m('span.fa.fa-bars')),
                m('form.form-inline.search-form.mr-2', [
                    m('.input-group', [
                        m('input.form-control[type=search][placeholder=Search...]', { onchange: (e) => state.setQuery(e.target.value), value: state.query }),
                        m('.input-group-append', [
                            m('button.btn.btn-outline-primary[type=submit]', { onclick: () => m.route.set("/search") }, m('i.fa.fa-search')),
                        ]),
                    ]),
                ]),
                m('ul.navbar-nav.ml-auto', [
                    m(curtime),
                    m(notifications),
                    m(profile),
                ])
            ])
        }
    }
}
