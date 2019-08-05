import m from 'mithril'

const Notice = {
    view(vnode) {
        return m('.reset-notice', [
            m('h1', "Password reset"),
            m('p', "A message with password reset instructions has been sent to your email."),
            m('p', "Please, click on the link inside to set a new password.")
        ])
    }
}

export default Notice