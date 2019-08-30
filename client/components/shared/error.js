import m from 'mithril'

const Error = {
    view(vnode) {
        let errors = vnode.attrs.errors
        return (Array.isArray(errors) && errors.length) ?
            m('.text-danger.validation-errors.mb-2', errors.map((e) => { return m('span', `${e} `) })) : null
    }
}

export default Error