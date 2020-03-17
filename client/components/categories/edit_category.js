import m from 'mithril'
import error from '../shared/error'
import {
    addSuccess
} from '../shared/notifications'
import {
    responseErrors,
} from '../../utils/helpers'
import service from '../../utils/service.js'

export default function Category() {
    let errors = [],
        category = {},
        isNew = true,
        loaded = false,

        setName = (name) => category.name = name,

        //requests
        newCategory = () => {
            category = {}
            loaded = true
        },

        editCategory = (id) =>
            service.getCategory(id)
                .then((result) => {
                    category = result;
                    loaded = true;
                }).catch((error) => errors = responseErrors(error)),

        create = () =>
            service.createCategory(category)
                .then((result) => {
                    addSuccess("Category created.")
                    m.route.set('/categories')
                })
                .catch((error) => errors = responseErrors(error)),

        update = () =>
            service.updateCategory(category.id, category)
                .then((result) => {
                    addSuccess("Category updated.")
                    m.route.set('/categories')
                })
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            isNew = (m.route.param('id') == undefined)
            if (isNew)
                newCategory()
            else
                editCategory(m.route.param('id'))
        },

        view(vnode) {
            return m(".category", (loaded) ? [
                m('h1.title', (isNew) ? 'New Category' : 'Edit Category'),
                m('.form-group', [
                    m('label', 'Name'),
                    m('input.form-control[type=text]', {
                        oncreate: (el) => el.dom.focus(),
                        oninput: (e) => setName(e.target.value),
                        placeholder: 'e.g. Busyness',
                        value: category.name
                    })
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
