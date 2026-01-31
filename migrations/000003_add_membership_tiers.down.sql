-- 000003_add_membership_tiers.down.sql
-- Rollback: 회원 등급 시스템 롤백

-- 외래키 제거
ALTER TABLE users DROP CONSTRAINT IF EXISTS fk_users_lotto_tier;

-- 인덱스 삭제
DROP INDEX IF EXISTS idx_membership_tiers_level;
DROP INDEX IF EXISTS idx_membership_tiers_code;
DROP INDEX IF EXISTS idx_users_lotto_tier;

-- membership_tiers 테이블 삭제
DROP TABLE IF EXISTS membership_tiers;
