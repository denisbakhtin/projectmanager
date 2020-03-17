import m from 'mithril';
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
            return m('nav.navbar.navbar-expand-lg.navbar-dark.navbar-bg', [
                m('.container', [
                    m('a.navbar-brand[href=#!/]', [
                        m('img', {
                            src: "/public/images/navbar_logo.png",
                            class: 'mr-2',
                            width: '23px',
                            height: '23px',
                        }),
                        site_name,
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
}
