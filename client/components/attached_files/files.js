import m from 'mithril'
import error from '../shared/error'
import {
    addSuccess
} from '../shared/notifications'
import {
    responseErrors
} from '../../utils/helpers'
import Auth from '../../utils/auth'
import service from '../../utils/service.js'

export default function AttachedFiles() {
    let files = [],
        errors = [],

        //requests
        upload = (e) => {
            let file = e.target.files[0]
            service.uploadFile(e.target.files[0])
                .then((result) => files.push(result))
                .catch((error) => errors = responseErrors(error))
        },
        create = () => {},
        remove = (index) => files.splice(index, 1)

    return {
        oninit(vnode) {
            files = []
            errors = []
            if (vnode.attrs.files && vnode.attrs.files.length > 0)
                files = vnode.attrs.files.slice(0)
        },
        onchange(vnode) {
            if (typeof vnode.attrs.onchange == 'function') vnode.attrs.onchange(files)
        },

        //TODO: use polymorphic gorm files relation (for tasks, projects, task_logs, etc)
        //show upload progress https://mithril.js.org/request.html#monitoring-progress
        view(vnode) {
            return m(".attached_files", [
                files.length > 0 ? files.map((file, index) => {
                    return m('span.badge.badge-light.mr-2', [
                        m('a[target=_blank][download]', {
                            href: file.url
                        }, file.name),
                        m('a[href=#]', {
                            onclick: () => {
                                remove(index);
                                vnode.attrs.onchange(vnode);
                                return false
                            }
                        }, m('i.fa.fa-times.text-danger.ml-2'))
                    ])
                }) : null,
                m('button.btn.btn-sm.btn-light', {
                    onclick: () => {
                        document.getElementById('attachment').click();
                        return false
                    }
                }, [
                    m('span', "Add"),
                    m('i.fa.fa-paperclip.ml-2')
                ]),
                m('input#attachment.hidden[type=file]', {
                    onchange: (e) => {
                        upload(e).then((result) => vnode.attrs.onchange(vnode))
                    }
                }),
            ])
        }
    }
}
