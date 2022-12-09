declare namespace API {
  type Response<T = any> = {
    data?: T;
    message?: string;
    statusCode?: number;
    success?: boolean;
    total?: number;
  };
  type PageQuery = {
    current?: number;
    pageSize?: number;
    query?: any;
    sorts?: any;
  };



  type SysDictionary = {
    create_at?: string;
    deleted_at?: string;
    id?: number;
    lable?: string;
    path?: string;
    pid?: number;
    update_at?: string;
    value?: string;
  };

  type SysOrgnization = {
    contact_addres?: string;
    contact_name?: string;
    contact_phone?: string;
    created_at?: string;
    deleted_at?: string;
    ext1?: string;
    ext10?: string;
    ext2?: string;
    ext3?: string;
    ext4?: string;
    ext5?: string;
    ext6?: string;
    ext7?: string;
    ext8?: string;
    ext9?: string;
    id?: number;
    name?: string;
    pid?: number;
    type_id?: number;
    updated_at?: string;
    /** 权限标识 */
    p_key?: string;
    /** 权限名称 */
    p_name?: string;
    path?: string
    children?: SysOrgnization[]
  };


  type OrgnizationType = {
    created_at?: string;
    id?: number;
    type_name?: string;
    updated_at?: string;
  };



  type SysRole = {
    created_at?: string;
    deleted_at?: string;
    id?: number;
    s_key?: string;
    name?: string;
    updated_at?: string;
    desc?: string
  };
}
