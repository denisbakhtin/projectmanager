import m from 'mithril'
import service from '../../utils/service.js'
import {
    addSuccess,
    addDanger
} from '../shared/notifications'
import {
    responseErrors
} from '../../utils/helpers'
import files from '../attached_files/files'
import error from '../shared/error'
import modal from '../shared/modal'

export default function EditComment() {
    let comment = {},
        task_id,
        is_solution,
        errors = [],
        onOk,
        onCancel,

        setContents = (contents) => comment.contents = contents,
        setFiles = (files) => comment.files = files,
        modalOnOk = () => (comment.id) ? update().then(onOk) : create().then(onOk),

        //requests
        create = () =>
            service.createComment(comment)
                .then((result) => addSuccess("New comment created."))
                .catch((error) => { errors = responseErrors(error); addDanger(errors.join('. ')) }),

        update = () =>
            service.updateComment(comment.id, comment)
                .then((result) => addSuccess("Comment updated."))
                .catch((error) => { errors = responseErrors(error); addDanger(errors.join('. ')) })

    return {
        oninit(vnode) {
            task_id = vnode.attrs.task_id
            is_solution = vnode.attrs.is_solution ?? false
            //atm solution is set only for new comments, may change in future without problem tho
            if (vnode.attrs.comment) {
                Object.assign(comment, vnode.attrs.comment) //copy object
            } else
                comment = { task_id: task_id, is_solution: is_solution }
            onOk = vnode.attrs.onOk ?? (() => null)
            onCancel = vnode.attrs.onCancel ?? (() => null)
        },
        view(vnode) {
            let body = m(".comment", [
                m('.form-group', [
                    m('label', "Contents (supports Markdown)"),
                    m('textarea.form-control', {
                        oncreate: (el) => el.dom.focus(),
                        oninput: (e) => setContents(e.target.value),
                        placeholder: "e.g. Good job!",
                        value: comment.contents
                    })
                ]),
                m('.form-group', [
                    m('label', "Attached files"),
                    m(files, {
                        files: comment.files,
                        onchange: setFiles
                    }),
                ]),
                m(error, { errors: errors }),
            ])

            return m(modal, {
                large: true,
                title: (comment.id) ? 'New comment' : 'Edit comment',
                body: body,
                extra_buttons: (is_solution) ? [
                    m('button[type=button].btn.btn-info.mr-auto', {
                        onclick: () => {
                            setContents('Done!')
                            modalOnOk()
                        }
                    }, "Done!"),
                ] : null,
                onOk: modalOnOk,
                onCancel: onCancel,
            })

        }
    }
}
