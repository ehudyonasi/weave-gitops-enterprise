import { Tab, Tabs, Theme } from '@material-ui/core';
import { createStyles, makeStyles } from '@material-ui/styles';
import { theme } from '@weaveworks/weave-gitops';
import styled from 'styled-components';
import { Dialog } from '@material-ui/core';
import { SubRouterTabs } from '@weaveworks/weave-gitops';

const { xs, small, medium, base, none } = theme.spacing;
const { primary10, primary, neutral30, neutralGray, white } = theme.colors;
const { small: smallSize } = theme.fontSizes;
const { monospace } = theme.fontFamilies;

export const useWorkspaceStyle = makeStyles(() =>
  createStyles({
    navigateBtn: {
      marginBottom: medium,
      marginRight: none,
      textTransform: 'uppercase',
    },
    filterIcon: {
      color: primary10,
      marginRight: small,
    },
    alertWrapper: {
      padding: base,
      margin: `0 ${base} ${base} ${base}`,
      borderRadius: '10px',
    },
    fullWidth: {
      width: '100%',
    },
    link: {
      color: primary,
      fontWeight: 600,
      whiteSpace: 'pre-line',
    },
    tabsWrapper: {
      'a[class*="MuiTab-root"]': {
        width: 'fit-content',
      },
    },
    YamlBtn: {
      width: '100%',
      display: 'flex',
      justifyContent: 'flex-end',
    },
  }),
);

export const CustomSubRouterTabs = styled(SubRouterTabs)(({}) => ({
  '.MuiTabs-root': {
    marginTop: medium,
    width: '100%',
    fontSize: smallSize,
    fontWeight: 600,
    minHeight: '32px',
    minWidth: '133px',
    opacity: 1,
    paddingLeft: '0 !important',
    paddingRight: '0 !important',
    span: {
      color: neutral30,
    },
    '.MuiTab-root': {
      minWidth: 'fit-content',
    },
    '.MuiTabs-indicator': {
      display: 'none !important',
    },
  },
  '.Mui-selected': {
    fontWeight: 700,
    background: `${primary}1A`,
    borderBottom: '3px solid #00b3ec',
    span: {
      color: primary10,
    },
  },
}));

export const DialogWrapper = styled(Dialog)`
  .MuiDialog-paper {
    border-radius: 10px;
  }
  .MuiDialogTitle-root {
    background: ${neutralGray};
    padding: ${medium};
    padding-bottom: ${small} ;
    p{
        font-weight: 600;
    }
    .MuiSvgIcon-root{
        color: ${neutral30};
    }
    .info{
        color: ${primary10} ;
        font-size: ${smallSize};
        font-weight: 500;
    }
  }
  .MuiDialogContent-root{
    &.customBackgroundColor{
      background: ${neutralGray} !important;
      padding:  ${none};
    }
    pre{
        background: ${white}!important;
        padding-left:${none} !important;
        span{
        font-family: ${monospace};
        font-size: ${smallSize};
        text-align: left !important;
        padding-right: ${none} !important;
        min-width: 27px !important;
    }
  }
    }
  }
`;

export const RulesListWrapper = styled.ul`
  list-style: none;
  margin-top: ${none} !important;
  padding-left: ${none} !important;
  li {
    background: ${white};
    margin-bottom: ${small};
    padding: ${small} ${medium};
    font-family: ${monospace};
    font-size: ${smallSize};
    label {
      margin-right: ${xs};
    }
  }
`;

export const ViewYamlBtn = styled.div``;
