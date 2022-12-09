import React, { useState, useEffect, ReactText } from 'react';
import { Upload, message, Button, UploadProps } from 'antd';
import { UploadFile as FileListItem } from 'antd/lib/upload/interface';
import { UploadOutlined } from '@ant-design/icons';
interface UploadFileProps extends UploadProps {
  action?: string;
  value?: string | string[];
  onChange?: (value: any) => void;
  limit?: number;
  isString?: boolean;
  disabled?: boolean;
}

interface File {
  uid: string;
  status: ReactText;
  url: string;
  name: string;
}

const UploadFile = function ({
  action,
  value,
  onChange,
  limit = 1,
  isString,
  disabled,
  ...rest
}: UploadFileProps) {
  const [hasEdit, $hasEdit] = useState<boolean>(false);
  const [fileList, $fileList] = useState<File[]>([]);
  // const { state } = useC2Mod('user')
  useEffect(() => {
    if (!hasEdit) {
      if (typeof value === 'string') {
        const cache = value.split(',').filter(x => x && x !== "");
        const fileList = cache.map((str, index) => ({
          uid: str,
          url: `/api/cms/file/get/${str}`,
          status: 'done',
          name: `附件${index}`,
        }));
        $fileList(fileList);
      } else if (Array.isArray(value)) {
        const fileList = value.map((str: string, index) => ({
          uid: str,
          url: `/api/file/get/${str}`,
          status: 'done',
          name: `附件${index}`,
        }));
        $fileList(fileList);
      }
    }
  }, [value]);

  const handleFileChange = function ({ fileList }: any) {

    if (limit && fileList.length > limit) {
      fileList = fileList.slice(1);
    }
    fileList = fileList.map((it: any) => {
      if (it.response && it.response.success) {
        return {
          uid: it.response.data.id,
          status: it.status,
          name: it.response.data.name,
          url: `/api/cms/file/get/${it.response.data.id}`,
        };
      }
      return it;
    });

    $fileList(fileList);
    $hasEdit(true);
    const doneList = fileList
      .filter((it: File) => it.status === 'done')
      .map((it: File) => it.uid);
    onChange && onChange(isString ? doneList.join(',') || ',' : doneList);
  };

  //console.log(fileList)

  return (
    <Upload
      action={action}
      fileList={fileList as FileListItem<any>[]}
      name="file"
      // headers={{ authorization: `Bearer ${storage.get('token')}` }}
      onChange={handleFileChange}
      disabled={disabled}
      {...rest}
      headers={{
        Authorization: `Bearer ${state['token']}`,
      }}

    >
      {!disabled && (
        <Button>
          <UploadOutlined /> 点击上传
        </Button>
      )}
    </Upload>
  );
};

export default UploadFile;
