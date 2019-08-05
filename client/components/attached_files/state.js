import m from 'mithril'
import {
    addSuccess
} from '../shared/notifications'
import {
    responseErrors
} from '../../utils/helpers'
import Auth from '../../utils/auth'

const state = {
    files: [],
    errors: [],
    //requests
    upload(e) {
        var file = e.target.files[0]
        var data = new FormData()
        data.append("upload", file)
        return m.request({
                method: "POST",
                url: `/api/upload/form`,
                body: data,
                headers: {
                    Authorization: Auth.authHeader()
                }
            })
            .then((result) => state.files.push(result))
            .catch((error) => state.errors = responseErrors(error))
    },
    create() {

    },
    remove(index) {
        state.files.splice(index, 1)
    },
}

export default state