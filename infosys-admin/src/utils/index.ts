import { truncate } from '@umijs/deps/compiled/lodash';
import _ from 'lodash'

export const isEmptyStr = (str?: string) => !str || str.trim().length === 0


export const deepArrayFind = (data: Array<any>, id: any) => _(data)
    .thru(function (coll) {
        return _.union(coll, _.map(coll, 'children') || []);
    })
    .flatten()
    .find({ id });

export const subStrDot = (str: string, length = 40) => {

    if (str && str.length > length) {
        return str.substr(0, length) + "..."
    }
    return str
}