import m from 'mithril'
import error from '../shared/error'
import state from './state'
import pusers from '../project_users/pusers'
import files from '../attached_files/files'
import status_state from '../statuses/state'

export default function form() {
    return [
        m('.form-group', [
            m('label', 'Project name'),
            m('input.form-control[type=text]', {
                oncreate: (el) => {
                    el.dom.focus()
                },
                oninput: (e) => state.setName(e.target.value),
                value: state.project.name
            })
        ]),
        m('.form-inline', [
            m('.form-group.mb-3.mr-4', [
                m('label.mr-2', 'Start date'),
                m('input.form-control[type=date]', {
                    oninput: (e) => state.setStartDate(e.target.value),
                    value: state.project.start_date
                })
            ]),
            m('.form-group.mb-3.mr-4', [
                m('label.mr-2', 'End date'),
                m('input.form-control[type=date]', {
                    oninput: (e) => state.setEndDate(e.target.value),
                    min: state.project.start_date,
                    value: state.project.end_date
                })
            ]),
            m('.form-group.mb-3.mr-4', [
                m('label.mr-2', 'Owner'),
                m('input.form-control[type=text][disabled]', {
                    value: state.getOwnerName()
                })
            ]),
            m('.form-group.mb-3', [
                m('label.mr-2', 'Status'),
                m('select.form-control', {
                        onchange: (e) => state.setStatusId(e.target.value),
                        value: state.project.status_id
                    }, status_state.statuses ?
                    status_state.statuses.map((status) => {
                        return m('option', {
                            value: status.id
                        }, status.name)
                    }) : null)
            ]),
        ]),
        m('.form-group', [
            m('label', "Assigned users"),
            m(pusers, {
                project_id: state.project.id,
                project_users: state.project.project_users,
                onchange: state.setProjectUsers
            }),
        ]),
        m('.form-group', [
            m('label', "Description"),
            m('textarea.form-control', {
                oninput: (e) => state.setDescription(e.target.value),
                value: state.project.description
            })
        ]),
        m('.form-group', [
            m('label', "Attached files"),
            m(files, {
                files: state.project.files,
                onchange: state.setFiles
            }),
        ]),
        m('.mb-2', m(error, {
            errors: state.errors
        })),
    ]
}