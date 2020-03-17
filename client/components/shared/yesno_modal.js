import m from 'mithril'
import modal from './modal'

export default function YesNoModal() {
    let onYes,
        onNo

    return {
        oninit(vnode) {
            onYes = vnode.attrs.onYes ?? (() => null)
            onNo = vnode.attrs.onNo ?? (() => null)
        },
        view(vnode) {
            return m(modal, {
                title: 'Confirmation is required',
                body: m('div', 'Are you sure?'),
                okText: 'Yes',
                cancelText: 'No',
                onOk: onYes,
                onCancel: onNo,
            })
        }
    }
}
