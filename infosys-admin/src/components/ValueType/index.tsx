import React from 'react';
import { ProRenderFieldPropsType } from '@ant-design/pro-provider'
import { ProColumns } from '@ant-design/pro-table'
import { toNumber } from 'lodash';
import { Input, TreeSelect } from 'antd';
import { deepArrayFind } from '@/utils';

export type ExProColumns<T = any> = ProColumns<T, 'icon' | 'number' | 'treeSelect'>

const ExValueType: Record<string, ProRenderFieldPropsType> = {


  number: {
    render: text => <div>{text}</div>,
    renderFormItem: (text, props) => (
      <Input value={props.fieldProps.value} type='number' onChange={e => props.fieldProps.onChange!(toNumber(e.target.value))} />
    ),
  },
  treeSelect: {
    render: (text, { fieldProps }) => {
      const filed = deepArrayFind(fieldProps.treeData, text)
      return <div>{filed && filed[fieldProps?.fieldNames?.label || 'name']}</div>
    },
    renderFormItem: (_, props) => <TreeSelect  {...props.fieldProps} />
  }
}

export default ExValueType;
