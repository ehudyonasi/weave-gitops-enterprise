import React, { FC } from 'react';
import { isEmpty } from 'lodash';
import styled from 'styled-components';
import { transparentize } from 'polished';
import { ReactComponent as BreadcrumbDivider } from '../assets/img/breadcrumb-divider.svg';
import { makeStyles, createStyles } from '@material-ui/core/styles';
import weaveTheme from 'weaveworks-ui-components/lib/theme';

interface Size {
  size?: 'small';
}

const Container = styled.div`
  align-items: flex-end;
  display: flex;
  font-size: ${20}px;
  height: 32px;
`;

export const Title = styled.div<Size>`
  margin-right: ${({ size }) =>
    size === 'small' ? weaveTheme.spacing.xxs : weaveTheme.spacing.xs};
  white-space: nowrap;
`;

const Link = styled.a`
  color: ${({ theme }) => theme.colors.black};
`;

export const Count = styled.div<Size>`
  background: ${({ size }) =>
    size === 'small'
      ? transparentize(0.5, weaveTheme.colors.purple100)
      : weaveTheme.colors.purple100};
  padding: 4px 8px;
  align-self: center;
  font-size: ${({ size }) =>
    size === 'small' ? weaveTheme.spacing.small : weaveTheme.fontSizes.normal};
  color: ${weaveTheme.colors.gray600};
  margin-left: ${weaveTheme.spacing.xs};
  border-radius: ${weaveTheme.borderRadius.soft};
`;

export interface Breadcrumb {
  label: string;
  url?: string;
  count?: number | null;
}

interface Props extends Size {
  path: Breadcrumb[];
}

const useStyles = makeStyles(() =>
  createStyles({
    path: {
      display: 'flex',
    },
    divider: {
      paddingLeft: weaveTheme.spacing.medium,
      paddingRight: weaveTheme.spacing.medium,
    },
  }),
);

export const Breadcrumbs: FC<Props> = ({ path, size }) => {
  const classes = useStyles();

  return (
    <Container>
      {path.map(({ label, url, count }, index) => (
        <div key={index} className={classes.path}>
          {index > 0 && (
            <div className={classes.divider}>
              <BreadcrumbDivider />
            </div>
          )}
          {isEmpty(url) ? (
            label
          ) : (
            <>
              <Title role="heading" size={size}>
                <Link href={url}>{label}</Link>
              </Title>
              {count !== null && (
                <Count className="section-header-count" size={size}>
                  {count}
                </Count>
              )}
            </>
          )}
        </div>
      ))}
    </Container>
  );
};