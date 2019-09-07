import m from 'mithril'

export default function Error() {
    return {
        view(vnode) {
            let errors = vnode.attrs.errors
            return (Array.isArray(errors) && errors.length) ?
                m('.text-danger.validation-errors.mb-2', errors.map((e) => {
                    return m('span', `${e} `)
                })) : null
        }
    }
}
