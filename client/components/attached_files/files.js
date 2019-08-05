import m from 'mithril'
import error from '../shared/error'
import state from './state'

const AttachedFiles = {
    oninit(vnode) {
        state.files = []
        state.errors = []
        if (vnode.attrs.files && vnode.attrs.files.length > 0) state.files = vnode.attrs.files.slice(0)
    },
    onchange(vnode) {
        if (typeof vnode.attrs.onchange == 'function') vnode.attrs.onchange(state.files)
    },
    
    //TODO: use polymorphic gorm files relation (for tasks, projects, task_logs, etc)
    //show upload progress https://mithril.js.org/request.html#monitoring-progress
    view(vnode) {
        let ui = vnode.state
        return m(".attached_files", [
            state.files.length > 0 ? state.files.map((file, index) => {
                return m('span.badge.badge-light.mr-2', [
                    m('a[target=_blank][download]', {href: file.url}, file.name),
                    m('a[href=#]', {onclick: () => {state.remove(index); ui.onchange(vnode); return false}}, m('i.fa.fa-times.text-danger.ml-2'))
                ])
            }) : null,
            m('button.btn.btn-sm.btn-light', {onclick: () => {document.getElementById('attachment').click();return false}}, [
                m('span', "Add"),
                m('i.fa.fa-paperclip.ml-2')
            ]),
            m('input#attachment.hidden[type=file]', {onchange: (e) => {state.upload(e).then((result) => ui.onchange(vnode))}}),
        ])
    }
}

export default AttachedFiles;