import m from 'mithril'

export default function Notice() {
  return {
    view(vnode) {
        return m('.activation-notice', [
            m('h1', "Activate your account"),
            m('p', "An activation message has been sent to your email."),
            m('p', "Please, click on the link inside to finish your registration.")
        ])
    }
  }
}
