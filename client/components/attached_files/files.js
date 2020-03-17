import m from 'mithril'
import error from '../shared/error'
import {
    addDanger
} from '../shared/notifications'
import {
    responseErrors
} from '../../utils/helpers'
import service from '../../utils/service.js'

export default function AttachedFiles() {
    let errors = [],
        readOnly = false,

        //requests 
        upload = (files, e) => {
            errors = []
            return service.uploadFile(e.target.files[0])
                .then((result) => files.push(result))
                .catch((error) => errors = responseErrors(error))
        },

        remove = (files, index) => files.splice(index, 1),
        onchange

    return {
        oninit(vnode) {
            readOnly = vnode.attrs.readOnly ?? false
            onchange = (typeof vnode.attrs.onchange == 'function') ? vnode.attrs.onchange : (() => null)
        },

        //show upload progress https://mithril.js.org/request.html#monitoring-progress
        view(vnode) {
            let files = (vnode.attrs.files) ? vnode.attrs.files.slice(0) : []
            let fullView = [
                files.map((file, index) => {
                    return m('span.badge.badge-light.mr-2', { key: file.id }, [
                        m('a[target=_blank]', {
                            href: file.url
                        }, file.name),
                        m('a[href=#]', {
                            onclick: () => {
                                remove(files, index);
                                onchange(files);
                                return false
                            }
                        }, m('i.fa.fa-times.text-danger.ml-2'))
                    ])
                }),
                m('button.btn.btn-sm.btn-default', {
                    onclick: () => {
                        document.getElementById('attachment').click();
                        return false
                    }
                }, [
                    m('span', "Add"),
                    m('i.fa.fa-paperclip.ml-2')
                ]),
                m(error, { errors: errors }),
                m('input#attachment.hidden[type=file]', {
                    onchange: (e) => upload(files, e).then((result) => onchange(files)),
                }),
            ],
                readOnlyView = (files.length > 0) ? m('.mt-2', [
                    m('span.fa.fa-paperclip.mr-2'),
                    files.map((file, index) => {
                        return m('span.badge.badge-light.mr-2', { key: file.id }, [
                            m('a[target=_blank]', {
                                href: file.url
                            }, file.name),
                        ])
                    })]) : null

            return m(".attached_files", readOnly ? readOnlyView : fullView)
        }
    }
}
