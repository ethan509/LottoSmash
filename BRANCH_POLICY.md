# 🔀 LottoSmash Branch Policy

이 문서는 LottoSmash 프로젝트의 Git branch 전략과 정책을 정의합니다.

## 📌 Branch Strategy: Trunk-Based Development

```
main (프로덕션, 항상 배포 가능)
  ↑ pull request (short-lived)
  |
feature/* (단기 - 3일 이내)
bugfix/* (단기 - 1일 이내)
hotfix/* (긴급 - 즉시)
```

**특징**:
- ✅ 빠른 배포 사이클 (일일 배포 가능)
- ✅ 간단한 브랜치 구조
- ✅ CI/CD 자동화 필수
- ✅ 코드 리뷰 철저
- ✅ 기능 플래그(Feature Flags) 권장

## 🌳 Branch 종류, Trunk)
- **목적**: 항상 배포 가능한 상태 유지
- **보호**: ✅ Branch Protection 활성화
- **규칙**:
  - PR 필수 (최소 1 승인)
  - 모든 CI/CD 체크 통과 필수
  - Squash merge만 허용
  - 단기 feature 브랜치에서만 merge
- **특징**: 모든 브랜치의 소스
- **배포**: 언제든 배포 가능 (수동 또는 자동)

### 2. **feature/*** (기능 개발 - 3일 이내)
- **명명규칙**: `feature/기능명` 또는 `feature/ISSUE-123-기능명`
- **생성**: `main`에서 생성
- **Merge**: `main`으로 PR 생성
- **수명**: 3일 이내 (최대)
- **예시**:
  - `feature/user-login`
  - `feature/ISSUE-42-auth-system`
  - `feature/email-verification`
- **주의**: 장기 개발은 feature flag 사용

### 3. **bugfix/*** (버그 수정 - 1일 이내)
- **명명규칙**: `bugfix/버그명` 또는 `bugfix/ISSUE-456-버그명`
- **생성**: `main`에서 생성
- **Merge**: `main`으로 PR 생성
- **우선순위**: 높음
- **수명**: 1일 이내
- **예시**:
  - `bugfix/login-error`
  - `bugfix/ISSUE-89-connection-timeout`

### 4. **hotfix/*** (긴급 버그 수정 - 즉시)
- **명명규칙**: `hotfix/버그명` 또는 `hotfix/ISSUE-999-버그명`
- **생성**: `main`에서 생성
- **Merge**: `main`으로 PR 생성 (긴급)
- **우선순위**: 최높음
- **수명**: 최소 (몇 시간 이내)
- **예시**:
  - `hotfix/security-patch`
  - `hotfix/ISSUE-999-data-loss`
- **태그**: merge 후 즉시 버전 태그 생성
- **Merge**: `main`으로 PR + tag 생성
- **목적**: 릴리스 전 버전 번프, 최종 테스트

## 📋 Commit Message Convention

**형식**: `<type>(<scope>): <subject>`

### Types
- `feat`: 새로운 기능
- `fix`: 버그 수정
- `docs`: 문서 추가/수정
- `style`: 코드 포맷, 세미콜론 등 (로직 변경 X)
- `refactor`: 로직 변경 없는 코드 개선
- `perf`: 성능 개선
- `test`: 테스트 추가/수정
- `chore`: 빌드, 의존성, CI/CD 설정 변경

### Scope (선택사항)
프로젝트의 영역을 명시:
- `auth`, `db`, `api`, `config`, `logger` 등

### Subject
- 명령조 사용 (Add, Fix, Update)
- 첫 글자 대문자
- 마침표 없음
- 50자 이내

### 예시
```
feat(auth): add email verification
fix(db): handle connection timeout
docs: update setup guide
refactor(api): simplify handler structure
perf(logger): optimize file rotation
```

## 🔄 PR (Pull Request) 프로세스

### 1. 브랜치 생성
```bash
git checkout develop
git pull origin develop
git checkout -b feature/새-기능-명
```

### 2. 개발 및 커밋
```bash
git add .main
git pull origin main
git checkout -b feature/새-기능-명
```

### 2. 개발 및 커밋 (짧은 사이클)
```bash
git add .
git commit -m "feat(scope): 변경사항 설명"
git push origin feature/새-기능-명

# 자주 push하기 (visibility)
```

### 3. PR 생성 (조기 생성 권장)
- GitHub에서 PR 생성 (진행 중이어도 OK, Draft 사용 가능)
- 템플릿 자동 로드 (.github/pull_request_template.md)
- 명확한 설명 작성
- 관련 이슈 연결 (`Closes #123`)

### 4. Code Review (빠른 피드백)
- 최소 1명의 승인 필수
- CODEOWNERS 자동 할당
- CI/CD 통과 필수 (테스트, 빌드, lint)
- 24시간 내 피드백
- 피드백에 즉시 대응

### 5. Merge (일일 배포)
- Squash merge (커밋 히스토리 정리)
- 자동 삭제 활성화 (feature 브랜치)
- **언제든 배포 가능 상태로 merge**

### 💡 핵심 원칙
- **짧은 수명**: feature는 최대 3일 이내
- **자주 merge**: 일일 1회 이상 권장
- **빠른 피드백**: 리뷰 지연 시간 최소화
- **테스트 필수**: 모든 PR은 CI/CD 통과 필수
- **기능 플래그**: 미완성 기능은 플래그로 숨기기
### 검증 항목
1. **Tests**: `go test ./...`
2. **Build**: `go build ./cmd/server`
3. **Lint**: golangci-lint (코드 품질)
4. **Docker Build**: Dockerfile 검증

### 실패 시 처리
- 모든 체크 통과 전까지 merge 불가
- 실패 원인 확인 및 (필수)
- ✅ Require pull requests before merging
  - Minimum 1 approval
  - Dismiss stale pull request approvals
- ✅ Require status checks to pass
  - test, build, lint
- ✅ Require branches to be up to date
- ✅ Require conversation resolution
- ✅ Include administrators
- ✅ Restrict who can push

**목표**: 항상 배포 가능 수명 |
|------|------|------|------|
| Feature | `feature/기능명` | `feature/user-auth` | 3일 이내 |
| Bug Fix | `bugfix/버그명` | `bugfix/login-error` | 1일 이내 |
| Hotfix | `hotfix/긴급버그` | `hotfix/security-patch` | 수 시간
## 📊 Naming Convention

| 타입 | 형식 |  (배포 마커)

### Release Tags
- 형식: `v1.2.3` (Semantic Versioning)
- 위치: `main` 브랜치에서만
- 생성 시점: 배포 후 또는 배포 전
- 자동 생성: GitHub Actions Release 워크플로우

### Semantic Versioning
- **Major (1.0.0)**: 호환되지 않는 변경
- **Minor (0.1.0)**: 새 기능 (하위 호환)
- **Patch (0.0.1)**: 버그 수정

### 예시
- `v1.0.0` - 첫 프로덕션 릴리스
- `v1.1.0` - 새 기능 추가
- `v1.1.1` - 긴급 버그 수정
- `v2.0.0` - 주요 변경사항ersioning)
- 위치: `main` 브랜치에서만
- 자동 생성: GitHub Actions Release 워크플로우

### 예시
- `v1.0.0` - 첫 메이저 릴리스
- `v1.1.0` - 마이너 버전 (새 기능)
- `v1.1.1` - 패치 버전 (버그 수정)

## ⚠️ 금지 사항
3일 이상 미merge)
❌ **merge commit** (squash merge만 사용)
❌ **테스트 없이 merge**
❌ **CI/CD 실패 상태로 merge**
❌ *🎯 Trunk-Based Development의 장점

| 항목 | Git Flow | Trunk-Based |
|------|----------|------------|
| 배포 주기 | 월/분기 | 일일 |
| 브랜치 복잡도 | 높음 | 낮음 |
| 병합 충돌 | 많음 | 적음 |
| 학습 곡선 | 가파름 | 완만함 |
| 롤백 | 어려움 | 쉬움 |

## 🛠️ 추가 권장사항

### Feature Flags (기능 플래그)
미완성 기능을 main에 merge하되 활성화하지 않기:
```go
if config.Features.EmailVerification {
    // 이메일 검증 코드
}
```

### Code Review 가이드
- 리뷰 시간: 24시간 이내
- 동기 토론: 면대면 또는 화상 회의
- 승인 기준: 코드 품질, 테스트 커버리지, 문서

### 배포 전략
- **Blue-Green**: 두 환경 동시 운영
- **Canary**: 점진적 배포
- **Feature Flags**: 런타임 토글

## 📚 참고자료

- [Trunk-Based Development](https://trunkbaseddevelopment.com/)
- [Conventional Commits](https://www.conventionalcommits.org/)
- [GitHub Branch Protection](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/managing-protected-branches)
- [CODEOWNERS](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/about-code-owners)
- [Feature Flags](https://www.atlassian.com/continuous-delivery/principles/feature-flag

- [Git Flow 설명](https://nvie.com/posts/a-successful-git-branching-model/)
- [Conventional Commits](https://www.conventionalcommits.org/)
- [GitHub Branch Protection](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/managing-protected-branches)
- [CODEOWNERS](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/about-code-owners)

## 📞 질문 및 피드백

정책에 대한 질문이나 개선사항은 이슈를 생성해주세요.
