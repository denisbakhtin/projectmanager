import m from 'mithril'
import notifications from './navbar_notifications'
import profile from './navbar_profile'
import curtime from './navbar_time'
import state from './state'
import service from '../../utils/service'

export default function Navbar() {
    let site_name = "Project Manager",

        getSettings = () =>
            service.getSettings()
                .then((result) => {
                    let settings = result.slice(0)
                    let name_setting = settings.find((el) => el.code === "site_name")
                    if (name_setting) site_name = name_setting.value
                })

    return {
        oninit(vnode) {
            getSettings()
        },

        view(vnode) {
            return m('nav.navbar.navbar-expand-md.navbar-light.bg-light.fixed-top', [
                m('a.navbar-brand[href=#!/]', [
                    m('img', {
                        src: "/public/images/navbar_logo.png",
                        class: 'mr-2',
                        width: '23px',
                        height: '23px',
                    }),
                    site_name,
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
