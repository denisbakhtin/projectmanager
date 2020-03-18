import m from 'mithril'
import comments_item from './comments_item'
import edit_comment_modal from './edit_comment_modal'

export default function Comments() {
    let task_id = {},
        showCommentsModal = false,
        onUpdate

    return {
        oninit(vnode) {
            task_id = vnode.attrs.task_id
            onUpdate = vnode.attrs.onUpdate ?? (() => null)
        },
        view(vnode) {
            let comments = vnode.attrs.comments
            return m('.comments', [
                m('.buttons.mb-4', [
                    m('button.btn.btn-primary.mr-2[type=button]', {
                        onclick: () => showCommentsModal = true
                    }, [
                        m('i.fa.fa-plus.mr-1'),
                        "Add Comment"
                    ]),
                ]),
                (comments && comments.length > 0) ?
                    m('ul.dashboard-box.box-list',
                        comments.map((comment) => m(comments_item, { key: comment.id, comment: comment, onUpdate: onUpdate }))
                    ) : m('p.text-muted', 'No comments yet.'),

                (showCommentsModal) ? m(edit_comment_modal, {
                    task_id: task_id,
                    onOk: () => { onUpdate(); showCommentsModal = false; },
                    onCancel: () => { showCommentsModal = false },
                }) : null,
            ])
        }
    }
}
