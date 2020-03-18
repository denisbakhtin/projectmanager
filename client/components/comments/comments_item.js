import m from 'mithril'
import MarkdownIt from 'markdown-it'
import {
    humanDate,
    responseErrors
} from '../../utils/helpers'
import {
    addDanger
} from '../shared/notifications'
import service from '../../utils/service.js'
import files from '../attached_files/files'
import edit_comment_modal from './edit_comment_modal'
import yesno_modal from '../shared/yesno_modal'

export default function CommentsItem() {
    let onUpdate,
        md,
        showCommentsModal = false,
        showRemoveModal = false,
        remove = (comment) =>
            service.deleteComment(comment.id)
                .then((result) => onUpdate())
                .catch((error) => addDanger(responseErrors(error).join('. ')))

    return {
        oninit(vnode) {
            onUpdate = vnode.attrs.onUpdate ?? (() => null)
            md = new MarkdownIt()
        },

        view(vnode) {
            let comment = vnode.attrs.comment
            return m('li', { class: comment.is_solution ? "solution-comment" : "" }, [
                m('.item-description', [
                    m('h3.item-title', [
                        m('.mr-2', m.trust(md.render(comment.contents))),
                        (comment.is_solution) ? m('span.badge.badge-success', "Solution") : null,
                    ]),
                    m('.dates', [
                        m('span.created-on.mr-3', [
                            m('span.fa.fa-calendar'),
                            m('span', 'Created on: '),
                            m('span', humanDate(comment.created_at)),
                        ]),
                        comment.updated_at > comment.created_at ? m('span.updated-on.mr-3', [
                            m('span.fa.fa-calendar'),
                            m('span', 'Updated on: '),
                            m('span', humanDate(comment.updated_at)),
                        ]) : null,
                    ]),
                ]),
                m(files, { files: comment.files, readOnly: true }),
                m('.buttons', [
                    m('button.btn.btn-default.btn-icon[type=button]', {
                        onclick: () => showCommentsModal = true
                    }, m('i.fa.fa-edit')),
                    m('button.btn.btn-default.btn-icon[type=button]', {
                        onclick: () => showRemoveModal = true
                    }, m('i.fa.fa-trash-o')),
                ]),

                (showCommentsModal) ? m(edit_comment_modal, {
                    task_id: comment.task_id,
                    comment: comment,
                    onOk: () => { showCommentsModal = false; onUpdate(); },
                    onCancel: () => { showCommentsModal = false },
                }) : null,

                (showRemoveModal) ? m(yesno_modal, {
                    onYes: () => { remove(comment); showRemoveModal = false },
                    onNo: () => showRemoveModal = false
                }) : null,
            ])
        }
    }
}
