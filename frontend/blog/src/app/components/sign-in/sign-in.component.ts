import {
  animate,
  state,
  style,
  transition,
  trigger,
} from "@angular/animations";
import { HttpClient } from "@angular/common/http";
import { Component, OnInit } from "@angular/core";
import {
  AbstractControl,
  FormBuilder,
  FormGroup,
  ValidatorFn,
  Validators,
} from "@angular/forms";

@Component({
  selector: "app-sign-in",
  animations: [
    trigger("formAnimation", [
      state("signup", style({ transform: "translateX(0%)" })),
      transition("signin => signup", [animate("1s 1.2s ease")]),
      state("signin", style({ transform: "translateX(100%)" })),
      transition("signup => signin", [animate("1s 1.2s ease")]),
    ]),
    trigger("leftPanelAnimation", [
      state(
        "signin",
        style({
          zIndex: 15,
          transform: "translateX(0)",
        })
      ),
      state(
        "signup",
        style({
          zIndex: 3,
          transform: "translateX(-150%)",
        })
      ),
      transition("signin => signup", [animate("1s 0.8s ease-in-out")]),
      transition("signup => signin", [animate("1s 0.8s ease-in-out")]),
    ]),
    trigger("rightPanelAnimation", [
      state(
        "signin",
        style({
          zIndex: 3,
          transform: "translateX(150%)",
        })
      ),
      state(
        "signup",
        style({
          zIndex: 15,
          transform: "translateX(0%)",
        })
      ),
      transition("signin => signup", [animate("1s 0.8s ease-in-out")]),
      transition("signup => signin", [animate("1s 0.8s ease-in-out")]),
    ]),
  ],
  templateUrl: "./sign-in.component.html",
  styleUrls: ["./sign-in.component.less"],
})
export class SignInComponent implements OnInit {
  signUpMode: boolean = false;
  signUpFormFlag: boolean = false;
  panelFlag: boolean = false;
  hiddenSignUpFormFlag: boolean = false;
  signInFormGroup: FormGroup;
  signUpFormGroup: FormGroup;
  signInData: any = {
    username: "",
    password: "",
  };
  signUpData: any = {
    username: "",
    email: "",
    tel: "",
    password: "",
  };
  signInFormErrorMsg: any = {
    username: "",
    password: "",
  };
  signUpFormErrorMsg: any = {
    username: "",
    password: "",
    tel: "",
    email: "",
  };
  validateMsg: any = {
    username: {
      required: "用户名不能为空",
      minlength: "用户名长度少于4个字符",
      maxlength: "用户名不能超过20个字符",
    },
    password: {
      required: "密码不能为空",
      minlength: "密码不能少于6个字符",
      maxlength: "密码不能超过20个字符",
    },
    tel: {
      required: "手机号不能为空",
      only: "手机号不正确",
    },
    email: {
      required: "邮箱不能为空",
      email: "邮箱不正确",
    },
  };

  constructor(private formBuilder: FormBuilder, private http: HttpClient) {
    this.signUpMode = false;
    this.signUpFormFlag = false;
    this.panelFlag = false;
    this.hiddenSignUpFormFlag = false;
  }

  ngOnInit(): void {
    this.signInFormGroup = this.formBuilder.group({
      username: [
        null,
        [
          Validators.required,
          Validators.minLength(4),
          Validators.maxLength(20),
        ],
      ],
      password: [
        null,
        [
          Validators.required,
          Validators.minLength(6),
          Validators.maxLength(20),
        ],
      ],
    });

    this.signInFormGroup.valueChanges.subscribe((data) =>
      this.onSignInFormValueChanged(data)
    );
    // 初始化错误信息
    this.onSignInFormValueChanged();

    this.signUpFormGroup = this.formBuilder.group({
      username: [
        null,
        [
          Validators.required,
          Validators.minLength(4),
          Validators.maxLength(20),
        ],
      ],
      password: [
        null,
        [
          Validators.required,
          Validators.minLength(6),
          Validators.maxLength(20),
        ],
      ],
      tel: [
        null,
        [
          Validators.required,
          this.validateRegExp(
            "only",
            /^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[01356789]\d{2}|4(?:0\d|1[0-2]|9\d))|9[189]\d{2}|6[567]\d{2}|4(?:[14]0\d{3}|[68]\d{4}|[579]\d{2}))\d{6}$/
          ),
        ],
      ],
      email: [null, [Validators.required, Validators.email]],
    });
    this.signUpFormGroup.valueChanges.subscribe((data) =>
      this.onSignUpFormValueChanged(data)
    );
    // 初始化错误信息
    this.onSignUpFormValueChanged();

    console.log(
      this.signUpMode,
      this.signUpFormFlag,
      this.panelFlag,
      this.hiddenSignUpFormFlag
    );
  }

  signIn() {
    console.log(this.signInData);
    this.http
      .post("api/api/v1/login", this.signInData)
      .toPromise()
      .then((data) => {
        console.log(data);
      })
      .catch((err) => {});
  }

  signUp() {
    console.log(this.signUpData);
    this.http
      .post("api/api/v1/register", this.signUpData)
      .toPromise()
      .then((data) => {
        console.log(data);
      })
      .catch((err) => {});
  }

  onSignUpFormAnimStart() {
    console.log(
      this.signUpMode,
      this.signUpFormFlag,
      this.panelFlag,
      this.hiddenSignUpFormFlag
    );
    this.panelFlag = !this.panelFlag;
    this.signUpMode = !this.signUpMode;
    setTimeout(() => {
      this.hiddenSignUpFormFlag = !this.hiddenSignUpFormFlag;
      this.signUpFormFlag = !this.signUpFormFlag;
    }, 1200);
  }

  onSignInFormValueChanged(data?: any) {
    // 如果表单不存在则返回
    if (!this.signInFormGroup) return;
    // 获取当前的表单
    const form = this.signInFormGroup;
    // 遍历错误消息对象
    for (const field in this.signInFormErrorMsg) {
      // 清空当前的错误消息
      this.signInFormErrorMsg[field] = "";
      // 获取当前表单的控件
      const control = form.get(field);
      // 当前表单存在此空间控件 && 此控件没有被修改 && 此控件验证不通过
      if (control && control.dirty && !control.valid) {
        // 获取验证不通过的控件名，为了获取更详细的不通过信息
        const messages = this.validateMsg[field];
        // 遍历当前控件的错误对象，获取到验证不通过的属性
        for (const key in control.errors) {
          // 把所有验证不通过项的说明文字拼接成错误消息
          this.signInFormErrorMsg[field] +=
            `<i class="fas fa-times-circle"></i>` + messages[key] + `\n`;
        }
      }
    }
  }

  onSignUpFormValueChanged(data?: any) {
    // 如果表单不存在则返回
    if (!this.signUpFormGroup) return;
    // 获取当前的表单
    const form = this.signUpFormGroup;
    // 遍历错误消息对象
    for (const field in this.signUpFormErrorMsg) {
      // 清空当前的错误消息
      this.signUpFormErrorMsg[field] = "";
      // 获取当前表单的控件
      const control = form.get(field);
      // 当前表单存在此空间控件 && 此控件没有被修改 && 此控件验证不通过
      if (control && control.dirty && !control.valid) {
        // 获取验证不通过的控件名，为了获取更详细的不通过信息
        const messages = this.validateMsg[field];
        // 遍历当前控件的错误对象，获取到验证不通过的属性
        for (const key in control.errors) {
          // 把所有验证不通过项的说明文字拼接成错误消息
          this.signUpFormErrorMsg[field] +=
            `<i class="fa fa-times-circle-o" aria-hidden="true"></i
          >` +
            messages[key] +
            `\n`;
        }
      }
    }
  }

  validateRegExp(type: string, regExp: RegExp): ValidatorFn {
    return (control: AbstractControl): { [key: string]: any } => {
      // 获取当前控件的内容
      const value = control.value;
      // 设置我们自定义的严重类型
      const res = {};
      res[type] = { value };
      // 如果验证通过则返回 null 否则返回一个对象（包含我们自定义的属性）
      return regExp.test(value) ? null : res;
    };
  }
}
