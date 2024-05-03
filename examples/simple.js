import {check} from 'k6';
import {MetainfoTest} from 'k6/x/stbb';

var test = MetainfoTest("cockroach://root@localhost:26257/metainfo?sslmode=disable")

export default function () {
    var res = test()
    check(res, {
        "no error": (res) => {
            return res == null
        }})

}
