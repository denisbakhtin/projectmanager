import m from 'mithril'
import error from '../shared/error'
import {
    addSuccess
} from '../shared/notifications'
import {
    responseErrors,
} from '../../utils/helpers'
import service from '../../utils/service.js'

export default function Page() {
    let errors = [],
        page = {},
        isNew = true,
        loaded = false,

        setName = (name) => page.name = name,
        setDescription = (description) => page.description = description,
        setMetaKeywords = (keywords) => page.meta_keywords = keywords,
        setMetaDescription = (description) => page.meta_description = description,
        setPublished = (value) => page.published = value,

        //requests
        newPage = () => {
            page = { published: true }
            loaded = true
        },

        editPage = (id) =>
            service.getPage(id)
                .then((result) => {
                    page = result;
                    loaded = true;
                }).catch((error) => errors = responseErrors(error)),

        create = () =>
            service.createPage(page)
                .then((result) => {
                    addSuccess("Page created.")
                    m.route.set('/pages')
                })
                .catch((error) => errors = responseErrors(error)),

        update = () =>
            service.updatePage(page.id, page)
                .then((result) => {
                    addSuccess("Page updated.")
                    m.route.set('/pages')
                })
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            isNew = (m.route.param('id') == undefined)
            if (isNew)
                newPage()
            else
                editPage(m.route.param('id'))
        },

        view(vnode) {
            return m(".page", (loaded) ? [
                m('h1.title', (isNew) ? 'New Page' : 'Edit Page'),
                m('.form-group', [
                    m('label', 'Name'),
                    m('input.form-control[type=text]', {
                        oncreate: (el) => el.dom.focus(),
                        oninput: (e) => setName(e.target.value),
                        placeholder: 'e.g. About Company',
                        value: page.name
                    })
                ]),
                m('.form-group', [
                    m('label', 'Contents (supports Markdown)'),
                    m('textarea.contents.form-control', {
                        oninput: (e) => setDescription(e.target.value),
                        value: page.description
                    })
                ]),
                m('.form-group', [
                    m('label', 'Meta Keywords'),
                    m('input.form-control[type=text]', {
                        oninput: (e) => setMetaKeywords(e.target.value),
                        value: page.meta_keywords
                    })
                ]),
                m('.form-group', [
                    m('label', 'Meta Description'),
                    m('input.form-control[type=text]', {
                        oninput: (e) => setMetaDescription(e.target.value),
                        value: page.meta_description
                    })
                ]),
                m('.form-group', [
                    m('input#published[type=checkbox]', {
                        onchange: (e) => setPublished(e.target.checked),
                        checked: page.published
                    }),
                    m('label.ml-2[for=published]', 'Published'),
                ]),
                m('.mb-2', m(error, { errors: errors })),
                m('.actions', [
                    m('button.btn.btn-primary.mr-2[type=button]', {
                        onclick: (isNew) ? create : update
                    }, [
                        m('i.fa.fa-check.mr-1'),
                        "Submit"
                    ]),
                    m('button.btn.btn-outline-secondary[type=button]', {
                        onclick: () => window.history.back()
                    }, "Cancel")
                ]),
            ] : m('Loading...'))
        }
    }
}
